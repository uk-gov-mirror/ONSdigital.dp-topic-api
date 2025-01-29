load("./scripts/utils/db.js");

/**
 * Makes deep copy of object
 * @param {object} obj - The object to copy.
 * @returns {object} - A true copy.
 */
function deepCopy(obj) {
  return JSON.parse(JSON.stringify(obj));
}

/**
 * Check if an ID is already used in the topics collection.
 * @param {string} id - The ID to check.
 * @returns {boolean} - True if the ID is used, otherwise false.
 */
function isUsedID(id) {
  return getTopicCollection().find({ id: id }).hasNext();
}

/**
 * Generate a random ID.
 * @returns {string} - The generated ID.
 */
function generateID() {
  let result = "";
  for (let i = 0; i < idSize; i++) {
    result += idAlphabet.charAt(Math.floor(Math.random() * idAlphabet.length));
  }
  return result;
}

/**
 * Generate an unused ID.
 * @returns {string} - The generated ID.
 */
function generateUnusedID() {
  let id;
  do {
    id = generateID();
  } while (isUsedID(id));

  return id;
}

/**
 * Generate a slug from a given title.
 * @param {string} title - The title to convert to a slug.
 * @returns {string} - The generated slug.
 */
function generateSlug(title) {
  return title.toLowerCase().replace(/[^a-z0-9]+/g, "");
}
