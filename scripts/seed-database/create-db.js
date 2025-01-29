load("./scripts/utils/config.js");
load("./scripts/utils/db.js");
load("./scripts/utils/utils.js");

function createCollections() {
  console.log("creating collections");
  if (!collectionExists(topicCollectionName)) {
    db.createCollection(topicCollectionName);
    getTopicCollection().createIndex({ id: 1 }, { name: "topics_id" });
    console.log(`${topicCollectionName} collection created`);
  } else {
    console.warn(
      `${topicCollectionName} collection already exists - not creating`
    );
  }

  if (!collectionExists(contentCollectionName)) {
    db.createCollection(contentCollectionName);
    getContentCollection().createIndex(
      { id: 1 },
      { name: "topics_content_id" }
    );
    console.log(`${contentCollectionName} collection created`);
  } else {
    console.warn(
      `${contentCollectionName} collection already exists - not creating`
    );
  }
}

function createRootTopic() {
  console.log("creating root topic");
  const rootTopic = {
    id: rootId,
    current: {
      id: rootId,
      state: "published",
      subtopics_ids: [],
    },
    next: {
      id: rootId,
      state: "published",
      subtopics_ids: [],
    },
  };
  insertTopic(rootTopic);
}
