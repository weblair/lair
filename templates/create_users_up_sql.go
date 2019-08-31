package templates

// CREATE_USERS_UP is the up-migration for the users table of a Gin project initialized with the --auth flag.
const CreateUsersUp = `CREATE TABLE users (
    id SERIAL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP,
    accept_after TIMESTAMP NOT NULL DEFAULT(NOW()),
    email VARCHAR(255) NOT NULL,
    email_verified BOOLEAN NOT NULL DEFAULT(FALSE),
    password BYTEA NOT NULL,
    public_id UUID NOT NULL,
    PRIMARY KEY(id),
    UNIQUE(email),
    UNIQUE(public_id)
);

CREATE INDEX users_public_id
ON users (public_id);

COMMENT ON TABLE users
IS 'User accounts with access to this API';

COMMENT ON COLUMN users.accept_after
IS 'On login, the API should only accept tokens that were issued after this date. '
   'This field is updated when the user changes their password.';

COMMENT ON COLUMN users.email
IS 'The email address that is associated with this account. This acts as the account username.';

COMMENT ON COLUMN users.email_verified
IS 'The email address for this user has been verified by the API.';

COMMENT ON COLUMN users.password
IS 'The hashed password to be used by the API to authenticate this user.';

COMMENT ON COLUMN users.public_id
IS 'The UUID that should be used when doing a public-facing lookup of this user (e.g. from a front-end URL).'
`
