# Go Utilities

## Common Encoding and Decoding
ecoding.go
- [x] `GetBasicAuthHeader(user string, password string) string `: RFC 4648 encoding string and return string serial.
       `DecodeBasicAuthHeader(header string) (string, string, error)`: RFC 4648 decoding string, vice verse as above function.
- [x] `encodeMd5(str string) string`: quickly gets the md5 code.
- [x] `EncodePassword(password string, salt string) string`:  encodes a password using PBKDF2.

