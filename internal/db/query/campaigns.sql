-- name: GetActiveCampaigns :many
SELECT * FROM campaigns WHERE status = 'ACTIVE';
