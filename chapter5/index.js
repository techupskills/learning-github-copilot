const express = require('express');
const app = express();
const port = 3000;

// Middleware to parse JSON
app.use(express.json());

// Import routes
const authController = require('./controllers/AuthController');
app.use('/auth', authController);

// Add a route for the root path
app.get('/', (req, res) => {
  res.send('Welcome to the Express.js application!');
});

app.listen(port, () => {
  console.log(`Server is running on http://localhost:${port}`);
});