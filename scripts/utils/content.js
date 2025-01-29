load('./scripts/utils/db.js');

/**
 * Find content by its id.
 * @param {string} id - The id of the content. This is not the mongo _id
 * @returns {object} - The found content object or null if not found.
 */
function findContentByID(id) {
    return getContentCollection().findOne({ id: id });
  }

/**
 * Create a content object for a topic.
 * @param {string} id - The ID of the topic.
 * @returns {object} - The created content object.
 */
function generateContent(id) {
  return {
    id: id,
    next: {
      state: "published",
    },
    current: {
      state: "published",
    },
  };
}

/**
 * Inserts a content document
 * @param {object} id - The new content document
 */
function insertContent(content) {
    console.log(`inserting content doc with id ${content.id}`);
    const contentWithIDExists = findContentByID(content.id);
    if (!contentWithIDExists) {
      if (cfg.insert) {
        getContentCollection().insertOne(content);
      }
    } else {
      console.warn(`content with id ${content.id} already exists`);
    }
  }


