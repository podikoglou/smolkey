package engine

// The maximum length that a key can have in bytes is 1024, since if a key was any
// bigger, it would be highly inefficient to look it up (probably) (you know what,
// I might be completely wrong, but if your key is bigger than 1024 bytes you're
// doing something wrong.)
const MAX_KEY_LENGTH = 1024

// Validates that a key (a string) is safe to put in the store. This returns
// `true` if it's safe to store, and `false' if it isn't.
func ValidateKey(key string) bool {
	return len(key) < MAX_KEY_LENGTH
}
