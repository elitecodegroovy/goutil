# Go Utilities

## Common Encoding and Decoding
ecoding.go

- [x] `GetBasicAuthHeader(user string, password string) string `: RFC 4648 encoding string and return string serial. 
       `DecodeBasicAuthHeader(header string) (string, string, error)`: RFC 4648 decoding string, vice verse as above function.
- [x] `encodeMd5(str string) string`: quickly gets the md5 code.
- [x] `encodeSha1(str string) string`: quickly gets the sha1 code.
- [x] `encodeSha256(str string) string`: quickly gets the sha256 code.
- [x] `encodeSha512(str string) string`: quickly gets the sha512 code.
- [x] `EncodePassword(password string, salt string) string`:  encodes a password using PBKDF2.
- [x] `Encrypt(payload []byte, secret string) ([]byte, error)`: encrypts a payload with a given secret.
       `Decrypt(payload []byte, secret string) ([]byte, error)`: decrypts a payload with a given secret.
       
filepath.go

- [ ] `encodeSha1(str string) string`: quickly gets the sha1 code.  
       

