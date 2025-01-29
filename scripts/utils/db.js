load("./scripts/utils/config.js");

/**
 * Gets the content collection
 * @returns {object} - The content collection
 */
function getContentCollection() {
  return db.getCollection(contentCollectionName);
}

/**
 * Gets the topic collection
 * @returns {object} - The topic collection
 */
function getTopicCollection() {
  return db.getCollection(topicCollectionName);
}

/**
 * Checks if a collection exists
 * @returns {boolean} - Does it exist
 */
function collectionExists(collectionName) {
  return db.getCollectionNames().includes(collectionName);
}
