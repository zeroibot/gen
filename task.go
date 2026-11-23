package main

import "uuid"

func generateUUIDv7() string {
	return uuid.NewV7().String()
}
