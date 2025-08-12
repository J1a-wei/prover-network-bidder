-- name: AddApp :exec
INSERT INTO app (
        app_id,
        img_url,
        registered
    )
VALUES (
        $1,
        $2,
        $3
    );


-- name: SelectMonitorBlock :one
SELECT * 
FROM monitor_block
WHERE event = $1;

-- name: UpsertMonitorBlock :exec
INSERT INTO monitor_block (event, block_num, block_idx, restart) 
VALUES ($1, $2, $3, $4) ON CONFLICT (event) DO
UPDATE
SET block_num = excluded.block_num,
    block_idx = excluded.block_idx,
    restart = restart;