const users = [
  { id: 1, username: 'user1', password: 'password1' },
  { id: 2, username: 'user2', password: 'password2' }
];

function authenticate(username, password) {
  return users.find(user => user.username === username && user.password === password);
}

module.exports = { authenticate };