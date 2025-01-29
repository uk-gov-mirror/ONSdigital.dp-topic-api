# Seeding database script

This folder contains a script used for seeding the `topics` database in mongodb.

Needs to have mongodb 4.4+ installed and running - this can be done via the dp-compose repository.

These scripts expect to be run from the root directory of this repo - there are make commands to do this:

```sh
    make database-seed
```

## Starting with blank database

Run:

`mongosh <mongo_url> index.js <options>`

This will:

- create the database and it's collections
- create a root topic for the topics to hang off
- insert Census topic and it's subtopics

The `<mongo_url>` part, if supplied, should look like:

- `<host>:<port>/<database_name>` for example: `localhost:27017/topics` (this is the default)
- If authentication is needed, use the format `mongodb://<username>:<password>@<host>:<port>`

Example of the (optional) `<options>` part:

- `--eval 'cfg={insert:false}'` (e.g. use for dry runs)
- `cfg` defaults to: `{insert: true}` (see utils/config.js)
- if you specify `cfg`, all missing options default to `false`

## Modifying the seed data

The seed data is contained with `./data.js` and can be modified there to add / remove topics.
