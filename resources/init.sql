

CREATE TABLE IF NOT EXISTS keyvalue (
    keyvalue_id serial PRIMARY KEY,
    key VARCHAR NOT NULL UNIQUE,
    value VARCHAR NOT NULL
)