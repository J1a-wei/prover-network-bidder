-- name: SaveApp :exec
INSERT INTO app (app_id, img_url, registered )
VALUES ($1, $2, $3) ON CONFLICT DO NOTHING;

-- name: AddProofRequest :exec
INSERT INTO proof_request (req_id, app_id, nonce, public_values_digest, input_data, input_url, max_fee, min_stake, deadline, created_at, processed)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11);

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

-- name: FindNotProcessedProofRequests :many
SELECT * FROM proof_request
WHERE processed = false;

-- name: UpdateRequestAsProcessed :exec
UPDATE proof_request
SET processed = true
WHERE req_id = $1;

-- name: AddBid :exec
INSERT INTO my_bid (req_id, my_fee, bid_nonce, should_reveal_after, should_reveal_before)
VALUES ($1, $2, $3, $4, $5);

-- name: FindToBeRevealedBid :many
SELECT * FROM my_bid
WHERE revealed = false AND ($1 BETWEEN should_reveal_after AND should_reveal_before);

-- name: UpdateBidAsRevealed :exec
UPDATE my_bid
SET revealed = true
WHERE req_id = $1;