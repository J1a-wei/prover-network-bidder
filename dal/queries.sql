-- name: AddApp :exec
INSERT INTO app (
        app_id,
        img_url
    )
VALUES (
        $1,
        $2
    );