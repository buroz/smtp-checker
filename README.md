```bash
$ ./bin/build.bash ./cmd/smtp-checker
```
```bash
$ ls -lah dist/
total 19M
-rwxr-xr-x  1 4.8M Mar  5 11:00 smtp-checker-darwin-amd64
-rwxr-xr-x  1 4.9M Mar  5 11:00 smtp-checker-linux-amd64
-rwxr-xr-x  1 4.3M Mar  5 11:00 smtp-checker-windows-386.exe
-rwxr-xr-x  1 4.8M Mar  5 11:00 smtp-checker-windows-amd64.exe
```

```bash
$ ./smtp-checker-linux-amd64 --help
Usage of ./dist/smtp-checker-linux-amd64:
  -host string
        Smtp host address
  -port int
        Smtp port number
  -receiver-email string
        Receiver's email
  -secure
        Smtp over TLS
  -sender-email string
        Sender's email
  -sender-password string
        Sender's password
```

```bash
./smtp-checker-linux-amd64 -host mail.larissahotels.com -port 465 -sender-email reservation1@larissahotels.com -sender-password !T0PSecreT! -receiver-email buroz@nethole.dev -secure
```
