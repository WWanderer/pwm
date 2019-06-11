# pwm -- Password manager written in Go
## functionalities
+ should be run using "pwd [pw database file]"
+ encrypts the password database file
+ CRUD the entries
+ interactive shell prompt (in the fdisk style) 
+ encoding using JSON
+ decryption on start/encryption on end

## json example 
```JSON
	{"Site":"zoomer.org", "Uname":"kevin", "Pw":"bazinga"}
```

## data structures
```Go
	type Entry struct {
		Site	string
		Uname	string
		Pw 		string
	}
```
## file organization
+ main.go
+ crud.go
	+ create(entry)
	+ read(entry)
	+ update(entry)
	+ delete(entry)
+ crypto.go
	+ encrypt(file)
	+ decrypt(file)
	+ verifyPassword(password)
+ util.go
	+ jsonParse ? (see: https://blog.golang.org/json-and-go)
	+ jsonWrite ?
	+ entryExists
	
## i/o
I will probably copy file content to memory, in an Entry slice, allowing for easier replacement
and deletion. The slice will be written to file after each operation.

## not yet decided/informed about
+ encryption and decryption
+ how to manage the master password
+ go project structure
+ capture terminal screen
