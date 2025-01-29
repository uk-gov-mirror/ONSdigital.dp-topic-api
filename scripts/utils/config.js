const topicDatabaseName = "topics";
const topicCollectionName = "topics";
const contentCollectionName = "content";

const idSize = 4;
const idAlphabet = "123456789";

const apiUrl = "http://localhost:25300/topics/";
const rootId = "topic_root";

const cfgDefault = {
  insert: true, // set to false to avoid inserts
};

if (typeof cfg == "undefined") {
  // Default configuration, can be overridden via command-line
  cfg = cfgDefault;
}
