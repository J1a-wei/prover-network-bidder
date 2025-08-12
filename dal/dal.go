package dal

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/cockroachdb/cockroach-go/v2/crdb"
	_ "github.com/lib/pq"
)

const (
	SqlUrlFmt = "postgresql://root@%s?sslmode=disable"
	TxTimeout = 5 * time.Second

	maxOpenConns = 20
)

var (
	bgCtx = context.Background()
)

type DAL struct {
	*Queries         // so all sqlc generated funcs are pass through
	raw      *sql.DB // underlying from sql.Open, don't use db as Queries also has db field
}

// dbUrl in format host:port/dbname, like localhost:26257/usualx_rewards
func NewDAL(dbUrl string) (*DAL, error) {
	sqldb, err := sql.Open("postgres", fmt.Sprintf(SqlUrlFmt, dbUrl))
	if err != nil {
		return nil, err
	}
	sqldb.SetMaxOpenConns(maxOpenConns)
	return &DAL{
		Queries: &Queries{sqldb},
		raw:     sqldb,
	}, nil
}

func (d *DAL) NewQueriesTx() (*Queries, *sql.Tx, error) {
	tx, err := d.raw.Begin()
	if err != nil {
		return nil, nil, err
	}
	return d.Queries.WithTx(tx), tx, nil
}

// minimal wrap, caller should do dtx := dal.Queries{tx} or dtx := dal.New(tx) in fn and call all dtx.XXX
// crdb.ExecuteTx handles retry based on fn return error
// fn must take care when wrapping errors returned from the database driver with additional context.
// ie. fn MUST use fmt.Errorf("%w", err) to ensure original db error is wrapped in return and available
// to crdb.ExecuteTx to decide what to do
// fn SHOULD avoid change any state that's not db. if can't avoid, state change must be idempotent,
// because fn may be run multiple times. non-db state change MUST be reverted by fn or upper level code if error
func (d *DAL) DoTx(fn func(tx *sql.Tx) error) error {
	ctx, cancel := context.WithTimeout(bgCtx, TxTimeout)
	defer cancel()
	return crdb.ExecuteTx(ctx, d.raw, nil, fn)
}

func (d *DAL) Close() error {
	if d.raw != nil {
		err := d.raw.Close()
		if err != nil {
			return err
		}
	}
	d.raw = nil
	d.Queries = nil
	return nil
}

// return epoch millisec
func Nowms() int64 {
	return time.Now().UnixNano() / 1e6
}

// if err is sql.ErrNoRows, return false, nil
func ChkQueryRow(err error) (bool, error) {
	found := false
	if err == nil {
		found = true
	} else if err == sql.ErrNoRows {
		err = nil
	}
	return found, err
}
