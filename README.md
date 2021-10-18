# Basic-Gin-Backend-Module


A basic Backend-System build with go, using the module viper, jwt, gorm, gin-gonic...etc.
There are the signup, login, me (get current user information) apis and DB CRUD (Create, Read, Update, Delete) for users.

**System Architecture**
- OS: Ubuntu 20.04.2 LTS with linux kernel 5.11.0-37-generic
- Go: go1.16.3 linux/amd64
- Docker: 20.10.5, build 55c4c88
- docker-compose: 1.25.0

**Note:**
- Postgres version must >= 12.0 because of the EXTENSION, pgcrypto.
