load("./scripts/utils/config.js");
load("./scripts/utils/db.js");

function wipe() {
  getTopicCollection().drop({});
  console.log(`${topicCollectionName} collection dropped`);
  getContentCollection().drop({});
  console.log(`${contentCollectionName} collection dropped`);
  db.dropDatabase();
  console.log(`${topicDatabaseName} db dropped`);
}

wipe();
