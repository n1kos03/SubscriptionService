-- +goose Up
-- +goose StatementBegin

INSERT INTO user_subscriptions (id, user_id, service_name, price, start_date)
VALUES
  (
    'c151f3d2-daf2-4c1a-8aa4-63bc22d53b82',
    '7639ca88-80df-4d91-ae20-78ac3431ee11',
    'Yandex Plus',
    400,
    '2025-07-01T00:00:00Z'
  ),
  (
    'a3f3b9c1-1a2b-4c8e-9f91-111111111111',
    '7639ca88-80df-4d91-ae20-78ac3431ee11',
    'Spotify',
    299,
    '2025-09-01T00:00:00Z'
  ),
  (
    'b4c2d7a8-9e4f-4e91-8d44-222222222222',
    '7639ca88-80df-4d91-ae20-78ac3431ee11',
    'Yandex Plus',
    400,
    '2025-08-01T00:00:00Z'
  ),
  (
    'd9e1f8c2-7a3b-4d22-9c55-333333333333',
    '11111111-2222-3333-4444-555555555555',
    'Netflix',
    599,
    '2025-07-01T00:00:00Z'
  );

-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin

DELETE FROM user_subscriptions
WHERE id IN (
  'c151f3d2-daf2-4c1a-8aa4-63bc22d53b82',
  'a3f3b9c1-1a2b-4c8e-9f91-111111111111',
  'b4c2d7a8-9e4f-4e91-8d44-222222222222',
  'd9e1f8c2-7a3b-4d22-9c55-333333333333'
);

-- +goose StatementEnd
