load("./scripts/seed-database/create-db.js");
load("./scripts/seed-database/insert-topics.js");

function init() {
  createCollections();
  createRootTopic();
  createTopics();
}

init();
