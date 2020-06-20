# go_ecrypt
Golang app the encrypt and decrypts text file 

The app does the following actions:
- requests the user a passphrase
- gets the 32byte hash of the passphrase
- encrtypts the content of the file using the hash
- writes the output to a new file
- reads the encrypted file, decrypts the contents, prints for user

App may be devided into two pieces one for encryption one for decryption.
