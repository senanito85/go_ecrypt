# go_ecrypt
Golang app the encrypt and decrypts text file 

- The app reads the given file 
- request the user a passphrase
- gets the 32byte hash of the passphrase
- encrtypt the content of the file using the hash
- Writes the output to a new file
- Reads the encrypted file, decrypts the contents, prints for user

App may be devided into two pieces one for encryption one for decryption.
