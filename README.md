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
       
       
## String Opts

- [x] `IsValidShortUID(uid string) bool`:  checks if short unique identifier contains valid characters.
- [x] `GenerateShortUID() string`:  generates a short unique identifier.

## time Opts

- [x] `GetCurrentDateISOStrDate() string`:  gets the current time with the format "2006-01-02".
      `GetISOStrDate(t time.Time) string `: does the same functionality.
- [x] `GetCurrentTimeISOStrTime() string`:  gets the current time with the format "2006-01-02 15:04:05".
      `GetISOStrTime(t time.Time) string`: does the same functionality.
- [x] `GetCurrentTimeNumberISOStrTime() string `:  gets the current time with the format "20060102150405".
      `GetISOStrTimeNumber(t time.Time) string `: does the same functionality.
- [x] `GetStrTime(t time.Time, format string) string `:  gets the current time with the provided format parameter.

## url Opts

- [x] `GetCurrentDateISOStrDate() string`:  gets the current time with the format "2006-01-02".


## File Opts
filepath.go

- [ ] `copyFiles(sourcePath string, targetPath string) error`: copy file from source directory to target directory.  
    
## IP Opts
ip_address.go

- [x] `ParseIPAddress(input string) string `: copy file from source directory to target directory.  
- [x] `SplitHostPortDefault(input, defaultHost, defaultPort string) (host string, port string)`:splits ip address/hostname string by host and port. Defaults used if no match found.

      
## JSON Opts


## XML Opts





       
## Log framework

Do import the package and then declare the `log` structure instance.

```
import (l "github.com/elitecodegroovy/goutil"
        "go.uber.org/zap"
)

var ( 
    log = l.GetLogger()
)

//...

func DoXxx(){
    log.Info("", zap.String("field name :", "value")
    //...
}
```

