const express = require('express');
const router = express.Router();
const userService = require('../services/UserService');

// Login route
router.post('/login', (req, res) => {
  const { username, password } = req.body;
  const user = userService.authenticate(username, password);
  if (user) {
    res.status(200).json({ message: 'Login successful', user });
  } else {
    res.status(401).json({ message: 'Invalid credentials' });
  }
});

module.exports = router;