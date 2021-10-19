# Basic-Gin-Backend-Module


A basic Backend-System build with go, using the module viper, jwt, gorm, gin-gonic...etc.
There are the signup, login, me (get current user information) apis and DB CRUD (Create, Read, Update, Delete) for users.

**System Architecture**
- OS: Ubuntu 20.04.2 LTS with linux kernel 5.11.0-37-generic
- Go: go1.16.3 linux/amd64
- Docker: 20.10.5, build 55c4c88
- Docker-Compose: 1.25.0
- Make: GNU Make 4.2.1, Built for x86_64-pc-linux-gnu

**Note:**
- Postgres version must >= 12.0 because of the EXTENSION, pgcrypto.

This example includes two runtimes, local and Docker. You can choose one to learn.

1. Build and Run at local
   1. make buildDB
   2. make build
   3. make run 
2. Build and Run on Docker
   1. make runOnContainer
