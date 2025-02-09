package main

import (
	"golang.org/x/crypto/bcrypt"
)

var correct_password string = "$2a$10$G6i6RQwxI68nd0xqoyZ1Lu5Ty7qsaBCScr4yr61DNe6zpOJda5uB." // This is the hash of "anirudhisreallycool"


func authenticate(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(correct_password), []byte(password))
	return err == nil
}

