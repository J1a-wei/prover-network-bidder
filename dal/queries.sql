-- name: SaveApp :exec
INSERT INTO app (app_id, img_url, register_status, register_error)
VALUES ($1, $2, $3, $4) ON CONFLICT DO NOTHING;

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
    restart = excluded.restart;

-- name: FindNotRegisteredApps :many
SELECT * FROM app
WHERE register_status = '';

-- name: UpdateAppAsRegisterSuccess :exec
UPDATE app
SET register_status = 'success'
WHERE app_id = $1;

-- name: UpdateAppAsRegisterFailed :exec
UPDATE app
SET register_status = 'failed', register_error = $2
WHERE app_id = $1;

-- name: ResetAppAsNotRegister :exec
UPDATE app
SET register_status = ''
WHERE app_id = $1;

-- name: FindNotProcessedProofRequests :many
SELECT p.* FROM proof_request p
INNER JOIN app a ON p.app_id = a.app_id
WHERE p.processed = false AND a.register_status = 'success';

-- name: UpdateRequestAsProcessed :exec
UPDATE proof_request
SET processed = true
WHERE req_id = $1;

-- name: AddBid :exec
INSERT INTO my_bid (req_id, my_fee, bid_nonce, should_reveal_after, should_reveal_before)
VALUES ($1, $2, $3, $4, $5);

-- name: FindToBeRevealedBid :many
SELECT * FROM my_bid
WHERE revealed = false AND ($1 > should_reveal_after AND $1 < should_reveal_before);

-- name: UpdateBidAsRevealed :exec
UPDATE my_bid
SET revealed = true
WHERE req_id = $1;

-- name: FindBidsWithoutResult :many
SELECT * FROM my_bid
WHERE bid_result = '' AND should_reveal_before < $1;

-- name: UpdateBidResult :exec
UPDATE my_bid
SET bid_result = $1
WHERE req_id = $2;

-- name: FindToBeProvedBids :many
SELECT b.*, p.app_id, p.input_data, p.input_url FROM my_bid b
INNER JOIN proof_request p
ON b.req_id = p.req_id
WHERE b.bid_result = 'success' AND b.proof_task_id = ''
AND p.deadline > $1;

-- name: UpdateBidProofTaskId :exec
UPDATE my_bid
SET proof_task_id = $1, proof_state = 'init'
WHERE req_id = $2;

-- name: FindBidsToQueryProvingResult :many
SELECT b.*, p.app_id FROM my_bid b
INNER JOIN proof_request p
ON b.req_id = p.req_id
WHERE b.proof_state = 'init' AND p.deadline > CAST(now() as BIGINT);

-- name: UpdateBidWithProof :exec
UPDATE my_bid
SET proof = $1, proof_state = 'generated'
WHERE req_id = $2;

-- name: FindBidsToSubmitProof :many
SELECT * FROM my_bid
WHERE proof_state = 'generated';

-- name: UpdateBidAsProofSubmitted :exec
UPDATE my_bid
SET proof_state = 'submitted', proof_submit_tx = $1
WHERE req_id = $2;