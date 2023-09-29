const mongoose = require("mongoose");

/**
 * Establishes a connection to the MongoDB database.
 * @param {string} uri - The MongoDB connection string.
 * @param {object} options - The MongoDB connection options.
 */
function connectToDatabase(uri, options) {
  return mongoose.connect(uri, options);
}

// Connect to the MongoDB database
const dbConnection = connectToDatabase(
  "mongodb://127.0.0.1:27017/securing-rest-apis-with-jwt",
  { useMongoClient: true }
);

// Handle successful connection
dbConnection.then(
  () => {
    console.log("Connected to MongoDB database");
  },
  (error) => {
    console.error("Error connecting to MongoDB database:", error);
  }
);
