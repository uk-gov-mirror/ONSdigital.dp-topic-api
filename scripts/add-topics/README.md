# Script for adding topics

This folder contains a script scripts used for adding items to the `topics` database in mongodb.

Needs to have mongodb 4.4+ installed and running - this can be done via the dp-compose repository.

These scripts expect to be run from the root directory of this repo - there are make commands to do this:

```sh
    make database-add
```

## Modifying the additional data

The added data is contained with `./data.js` and can be modified there to add / remove topics.

You will need to update the topic id in the data to match whatever topics you wish to add to.
