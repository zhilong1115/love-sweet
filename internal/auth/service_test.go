package auth

import (
	"testing"
)

func TestUserService_CreateUser(t *testing.T) {
	db := SetupTestDB()
	defer db.Close()

	s := NewUserService(db)

	// Test create user
	err := s.CreateUser("testuser", "testpassword")
	if err != nil {
		t.Fatal(err)
	}

	// Test create duplicate user
	err = s.CreateUser("testuser", "testpassword")
	if err == nil {
		t.Error("Expected error but got nil")
	}
}

func TestUserService_VerifyUser(t *testing.T) {
	db := SetupTestDB()
	defer db.Close()

	s := NewUserService(db)

	// Create test user
	s.CreateUser("testuser", "testpassword")

	// Test verify correct password
	ok, err := s.VerifyUser("testuser", "testpassword")
	if err != nil {
		t.Fatal(err)
	}
	if !ok {
		t.Error("Expected true but got false")
	}

	// Test verify incorrect password
	ok, err = s.VerifyUser("testuser", "wrongpassword")
	if err != nil {
		t.Fatal(err)
	}
	if ok {
		t.Error("Expected false but got true")
	}

	// Test verify non-existent user
	ok, err = s.VerifyUser("nonexistentuser", "testpassword")
	if err != nil {
		t.Fatal(err)
	}
	if ok {
		t.Error("Expected false but got true")
	}
}
