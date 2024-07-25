// insert-topics.js
//
// Script to create topics and subtopics in MongoDB

const topicsCollection = 'topics';  // Collection name for topics
const contentCollection = 'content';  // Collection name for content
const idSize = 4;  // Length of the generated ID
const idAlphabet = '123456789';  // Characters used for generating the ID
const apiUrl = "http://localhost:25300/topics/";  // Base URL for the API

if (typeof(cfg) == "undefined") {
    // Default configuration, can be overridden via command-line
    cfg = {
        verbose:  false,  // Display the new documents
        insert:   true    // Set value to false to avoid actual inserts
    }
}

// Array of topics to be added
const newTopics = [
    {
        title: "New Topic 1",
        description: "Description for New Topic 1",
        parentSlug: "parent_topic_slug_1",  // Slug of the parent topic for this topic
        subtopics: [
            { title: "Example Subtopic 1", description: "Description for Example Subtopic 1" },
            { title: "Example Subtopic 2", description: "Description for Example Subtopic 2" },
        ]
    },
    {
        title: "New Topic 2",
        description: "Description for New Topic 2",
        parentSlug: "parent_topic_slug_2",  // Slug of the parent topic for this topic
        subtopics: [
            { title: "Example Subtopic 3", description: "Description for Example Subtopic 3" },
            { title: "Example Subtopic 4", description: "Description for Example Subtopic 4" },
        ]
    }
];

/**
 * Check if an ID is already used in the topics collection.
 * @param {string} id - The ID to check.
 * @returns {boolean} - True if the ID is used, otherwise false.
 */
function isUsedId(id) {
    return db.getCollection(topicsCollection).find({ id: id }).hasNext();
}

/**
 * Generate a random ID.
 * @returns {string} - The generated ID.
 */
function makeId() {
    var result = '';
    for (var i = 0; i < idSize; i++) {
        result += idAlphabet.charAt(Math.floor(Math.random() * idAlphabet.length));
    }
    return result;
}

/**
 * Generate a slug from a given title.
 * @param {string} title - The title to convert to a slug.
 * @returns {string} - The generated slug.
 */
function generateSlug(title) {
    return title.toLowerCase().replace(/[^a-z0-9]+/g, '');
}

/**
 * Create a topic object.
 * @param {string} title - The title of the topic.
 * @param {string} description - The description of the topic.
 * @returns {object} - The created topic object.
 */
function createTopic(title, description) {
    var id;
    do {
        id = makeId();
    } while (isUsedId(id));

    var slug = generateSlug(title);

    var topic = {
        id: id,
        next: {
            id: id,
            slug: slug,
            description: description,
            title: title,
            state: "published",
            links: {
                self: {
                    href: apiUrl + id,
                    id: id
                },
                content: {
                    href: apiUrl + id + "/content"
                }
            }
        },
        current: {
            id: id,
            slug: slug,
            description: description,
            title: title,
            state: "published",
            links: {
                self: {
                    href: apiUrl + id,
                    id: id
                },
                content: {
                    href: apiUrl + id + "/content"
                }
            }
        }
    };

    return topic;
}

/**
 * Create a content object for a topic.
 * @param {string} id - The ID of the topic.
 * @returns {object} - The created content object.
 */
function createContent(id) {
    return {
        id: id,
        next: {
            state: "published"
        },
        current: {
            state: "published"
        }
    }
}

/**
 * Find a topic by its slug.
 * @param {string} slug - The slug of the topic.
 * @returns {object} - The found topic object or null if not found.
 */
function findTopicBySlug(slug) {
    var topicCursor = db.getCollection(topicsCollection).find({ 'current.slug': slug });
    return topicCursor.hasNext() ? topicCursor.next() : null;
}

// Create and add each new topic and its subtopics
newTopics.forEach(newTopicDef => {
    var parentTopic = findTopicBySlug(newTopicDef.parentSlug);
    if (!parentTopic) {
        print("Error: Couldn't find the parent topic with slug: " + newTopicDef.parentSlug);
        return;
    }

    var newTopic = createTopic(newTopicDef.title, newTopicDef.description);
    newTopic.next.links.subtopics = {
        href: apiUrl + newTopic.id + "/subtopics",
    };
    newTopic.current.links.subtopics = {
        href: apiUrl + newTopic.id + "/subtopics",
    };
    newTopic.next.subtopics_ids = [];
    newTopic.current.subtopics_ids = [];

    newTopicDef.subtopics.forEach(subtopicDef => {
        var subtopic = createTopic(subtopicDef.title, subtopicDef.description);
        var content = createContent(subtopic.id);

        if (cfg.verbose) {
            print("New subtopic document");
            print(JSON.stringify(subtopic));
            print("New content document");
            print(JSON.stringify(content));
        }

        if (cfg.insert) {
            db.getCollection(topicsCollection).insertOne(subtopic);
            db.getCollection(contentCollection).insertOne(content);
        }

        // Add subtopic to the new topic
        newTopic.next.subtopics_ids.push(subtopic.id);
        newTopic.current.subtopics_ids.push(subtopic.id);
    });

    var newContent = createContent(newTopic.id);
    if (cfg.verbose) {
        print("New topic");
        print(JSON.stringify(newTopic));
        print("New content");
        print(JSON.stringify(newContent));
    }

    if (cfg.insert) {
        // Insert the new topic and its content into the database
        db.getCollection(topicsCollection).insertOne(newTopic);
        db.getCollection(contentCollection).insertOne(newContent);

        // Add the new topic to the subtopics of the parent topic
        parentTopic.next.subtopics_ids.push(newTopic.id);
        parentTopic.current.subtopics_ids.push(newTopic.id);

        // Update the parent topic in the database
        db.getCollection(topicsCollection).updateOne({ 'current.slug': newTopicDef.parentSlug }, { $set: parentTopic });
    }
});