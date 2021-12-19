package database

const DatabaseSchema = `
CREATE TABLE IF NOT EXISTS blog (
    title varchar(200),
    description varchar(200),
    author varchar(250)
);
`
