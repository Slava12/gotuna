package doubles

import "github.com/Slava12/gotuna"

// MemUser1 is a sample in-memory user.
var MemUser1 = gotuna.InMemoryUser{
	ID:       "123",
	Email:    "john@example.com",
	Name:     "John",
	Password: "pass123",
}

// MemUser2 is a sample in-memory user.
var MemUser2 = gotuna.InMemoryUser{
	ID:       "456",
	Email:    "bob@example.com",
	Name:     "Bob",
	Password: "bobby5",
}

// NewUserRepositoryStub returns a new user repository with two sample users.
// This is a sample UserRepository implementation with users stored in-memory.
func NewUserRepositoryStub() gotuna.UserRepository {
	return gotuna.NewInMemoryUserRepository([]gotuna.InMemoryUser{
		MemUser1,
		MemUser2,
	})
}
