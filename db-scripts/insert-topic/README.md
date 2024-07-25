# insert-topics

This utility inserts topics and their subtopics into a MongoDB collection. Each new topic can be added under a different specified parent topic identified by its `slug`.

## How to run the utility

Run

```
mongo <mongo_url> <options> insert-topics.js
```

The `<mongo_url>` part should look like:

* `<host>:<port>/<database>` for example: `localhost:27017/topics`
* If authentication is needed, use the format `mongodb://<username>:<password>@<host>:<port>/<database>`
* In the above, `<database>` indicates the database to be modified

Example of the (optional) `<options>` argument:

`--eval 'cfg={verbose:true}'` (e.g. use for debugging)
cfg defaults to: {verbose:false, insert: true}
If you specify cfg, all missing options default to false

It is recommended to perform a dry run and check the result looks as expected:

```
mongo localhost:27017/topics --eval 'cfg={verbose:true, insert:false}' insert-topics.js
```

### Parameters to Configure

Edit the script to set the following variables according to your requirements:

* `newTopics`: An array of objects representing the new topics. Each object should have a `title`, `description`, `parentSlug`, and an array of `subtopics`. Each subtopic should have a `title` and `description`.


### TLS Configuration

When connecting to a TLS-enabled DocumentDB cluster (sandbox or prod), you'll need to add the following options:

* `--tls`
* `--tlsCAFile=<pem>` where `<pem>` is the path to the Certificate Authority .pem file

For example:

```
mongo mongodb://$MONGO_USER:$MONGO_PASS@$MONGO_HOST/topics --tls --tlsCAFile=./cert.pem insert-topics.js
```

### Example Run Usage

```
# Perform a dry run with verbose output
mongo localhost:27017/topics --eval 'cfg={verbose:true, insert:false}' insert-topics.js

# Insert topics and subtopics into the database
mongo localhost:27017/topics --eval 'cfg={verbose:false, insert:true}' insert-topics.js
```

### Notes

* Ensure the parent topics with the specified `parentSlug`  values exist in the database before running the script.
* The script generates unique slugs for each topic and subtopic based on their titles.
