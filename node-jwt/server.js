const express = require("express");
const app = express();
const port = process.env.PORT || 3000;

/**
 * Starts the Express server.
 * @param {number} port - The port number to start the server on.
 * @returns {http.Server} - The server instance.
 */
function startServer(port) {
  return app.listen(port, () => {
    console.log(`Express server listening on port ${port}`);
  });
}

// Start the server
const server = startServer(port);
