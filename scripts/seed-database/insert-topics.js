load("./scripts/seed-database/data.js");
load("./scripts/utils/topic.js")

function createTopics() {
  console.log("creating seed topics");
  seedTopics.forEach((topicData) =>
    createTopic(rootId, topicData, generateUnusedID())
  );
}

