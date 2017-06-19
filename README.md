# Golang on Ubuntu

#### SET UP
 * download : [https://golang.org/dl/](https://golang.org/dl/)
 * cd ~/Download run `sudo tar -C /usr/local -xzf go1.8.3.linux-amd64.tar.gz`
 * `sudo ln -s /usr/local/go/bin/go /usr/bin/go`
 * `go version`

#### Go set path variables
 * create Workspace at Home `mkdir -p Workspace/go`
 * `vi .profile`
   * export GOPATH=$HOME/Workspace/go
   * export PATH=$HOME/Workspace/go:$PATH
 * `go env`
