# pwm -- Password manager written in Go
## functionalities
+ should be run using "pwd [pw database file]"
+ encrypts entries before writing to file
+ CRUD the entries
+ interactive shell prompt (in the fdisk style) 
+ encoding using JSON
+ decryption to memory - file stays encrypted

## json example 
```JSON
	{"Site":"zoomer.org", "Uname":"kevin", "Pw":"bazinga"}
```

## data structures
```Go
	type Entry struct {
		Site	string
		Uname	string
		Pw 	string
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
	+ encrypt([]Entry)
	+ decrypt([]Entry)
	+ verifyPassword(password)
	+ genPW(length)
+ util.go
	+ loadFile
	+ writeToFile
	+ isNil(entry)
	+ createEntry
	
## not yet decided/informed about
+ ~~encryption and decryption~~ aes
+ ~~how to manage the master password~~ prolly shasum of a master passphrase
+ capture terminal screen
