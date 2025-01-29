load("./scripts/add-topics/data.js");
load("./scripts/utils/topic.js");
load("./scripts/utils/utils.js");


function addTopics() {
    console.log("creating new topics");
    addTopicsData.forEach((topicData) =>
      createTopic(topicData.parentID, topicData, generateUnusedID())
    );
  }
  
  addTopics();
