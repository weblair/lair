package templates

// CREATE_USERS_DOWN is the down-migration for the users table of a Gin project initialized with the --auth flag.
const CreateUsersDown = `DROP INDEX users_public_id;

DROP TABLE users;`
