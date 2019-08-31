package templates

const SeedDevelopmentYml = `users:
  - created_at: '2019-01-01T00:00'
    updated_at: '2019-01-01T00:00'
    accept_after: '2019-01-01T00:00'
    email: 'admin@example.com'
    email_verified: true
    password: '$PASSWORD(BADPASSWORD)"'
    public_id: '322a0ffb-3a98-43c4-852b-af5f9c6fa5e6'
  - created_at: '2019-01-01T00:00'
    updated_at: '2019-01-01T00:00'
    accept_after: '2019-01-01T00:00'
    email: 'user1@example.com'
    email_verified: true
    password: '$PASSWORD(BADPASSWORD)"'
    public_id: '034e3126-b612-437b-bff2-b3f114d157b3'
  - created_at: '2019-01-01T00:00'
    updated_at: '2019-01-01T00:00'
    accept_after: '2019-01-01T00:00'
    email: 'user2@example.com'
    email_verified: false
    password: '$PASSWORD(BADPASSWORD)"'
    public_id: '62e407a8-36a3-4f40-a469-eaf2a91cc6df'
  - created_at: '2019-01-01T00:00'
    updated_at: '2019-01-01T00:00'
    deleted_at: '2019-03-01T00:00'
    accept_after: '2019-01-01T00:00'
    email: 'user3@example.com'
    email_verified: true
    password: '$PASSWORD(BADPASSWORD)"'
    public_id: 'a138d351-210c-4d47-b9ae-e8e1ea2fac73'
`
