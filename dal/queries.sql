-- name: SaveApp :exec
INSERT INTO app (app_id, img_url, registered )
VALUES ($1, $2, $3) ON CONFLICT DO NOTHING;

-- name: AddProofRequest :exec
INSERT INTO proof_request (req_id, app_id, nonce, public_values_digest, input_data, input_url, max_fee, min_stake, deadline, created_at)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10);

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

-- name: FindNotRegisteredApps :many
SELECT * FROM app
WHERE registered = false;

-- name: UpdateAppAsRegistered :exec
UPDATE app
SET registered = true
WHERE app_id = $1;