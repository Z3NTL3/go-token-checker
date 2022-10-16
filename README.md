# go-token-checker
A Go script to check Discord tokens. Empowered by Goroutines.

# PROXY SUPPORT SOON
![layer](https://user-images.githubusercontent.com/48758770/196030162-0629bd03-9513-4a58-898f-d01d047b47f5.png)

```
I will add a proxied pool to this tool SOON. Currently I am busy with school.
```

```
# Compiling is not required but if you do not trust the executable , compile the SRC your self. 
# Use the instructions mentioned below
```

# How to run
On your terminal:
```./token-checker.exe tokenfile.txt```

### How to compile
One-time run:
```go run .```

After that run one-time:
``go build`` to compile everything so you can have an executable file.

Now your installation is done, just run the executable and there u go


### How to install Go
Install Go ``minimum Go version: 1.19``

Navigate to ``https://go.dev/dl/`` install the one you need compabitle with your OS.

```
# Installation
rm -rf /usr/local/go && tar -C /usr/local -xzf go1.19.2.linux-amd64.tar.gz 

// you may need to run the above command as sudo, if you do change alse the $HOME/.profile of root with source $HOME/.profile

export PATH=$PATH:/usr/local/go/bin
source $HOME/.profile

go version // if you get output it works. Do not forget to follow the last 2 steps on differents users on your machine
```
