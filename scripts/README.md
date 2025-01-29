# Scripts

This folder contains a scripts used for manipulating the `topics` database in mongodb.

Needs to have mongodb 4.4+ installed and running - this can be done via the dp-compose repository.

These scripts expect to be run from the root directory of this repo - there are make commands to do this:

```sh
    make database-seed # seeds a blank dataabase with test data
    make database-wipe # wipes the database
    make database-add # adds topics to an existing topic structure
```
