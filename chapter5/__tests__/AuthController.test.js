const AuthController = require('../path/to/AuthController'); // Update with the correct path

test('should register a user', () => {
	const result = AuthController.register('testUser', 'testPassword');
	expect(result).toBe(true);
});

test('should login a user', () => {
	const result = AuthController.login('testUser', 'testPassword');
	expect(result).toBe(true);
});

test('should fail to login with incorrect password', () => {
	const result = AuthController.login('testUser', 'wrongPassword');
	expect(result).toBe(false);
});