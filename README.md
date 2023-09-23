[![Go Reference](https://pkg.go.dev/badge/github.com/whonion/go-client-faucet-request.svg)](https://pkg.go.dev/github.com/whonion/go-client-faucet-request) [![Go version][go-badge]][go-url] [![go-report][go-report-badge]][go-report-url] [![Lint][lint-badge]][lint-url] [![Test][test-badge]][test-url] [![Build][build-badge]][build-url] [![Makefile][makefile-badge]][makefile-url]
# Faucet-client for send multi-request to server using proxy</br>
Implementation with go-routines and without
## Package description:
 - main.go - main package to run or build (sequential proxy request sending)
 - go-routine.go - multi-send requests with go-routine(in testing stage)
 - proxy.go - package for getting a proxy list from a remote site(in testing stage)
 ## Description of required files:
 - proxy.txt - proxy list file in the protocol://ip:port@login@pass format
 - createwallets.sh  - batch generator of wallets with output of 'addresses.txt' file
 - removewallets.sh - batch remover of wallets
 - addresses.txt - wallet's list for sending curl requests to target faucet's url
 - agents.txt - list of user-agents

## How to run with shell(without build) :
```sh
ver="1.20" && \
wget "https://golang.org/dl/go$ver.linux-amd64.tar.gz" && \
sudo rm -rf /usr/local/go && \
sudo tar -C /usr/local -xzf "go$ver.linux-amd64.tar.gz" && \
rm "go$ver.linux-amd64.tar.gz" && \
echo "export PATH=$PATH:/usr/local/go/bin:$HOME/go/bin" >> $HOME/.bash_profile && \
source $HOME/.bash_profile && \
go version \
sudo apt-get install tmux \
tmux new -s go_faucet
```
```sh
git clone https://github.com/whonion/go-client-faucet-request.git \
cd go-client-faucet-request \
#make the file 'createwallets.sh ' executable
chmod +x createwallets.sh \
# generate wallet addresses for the Tendermint node and out them in 'addresses.txt'
./createwallets.sh \
# add your proxy to proxy.txt. The number of proxies must match the number of addresses
go run  main.go
```
```sh
# or run with goroutine (recomended use it)
go run  go-routine.go
```
[go-badge]: https://img.shields.io/badge/go-1.20-blue.svg
[go-url]: https://go.dev

[go-report-badge]: https://goreportcard.com/badge/github.com/whonion/go-client-faucet-request
[go-report-url]: https://goreportcard.com/report/github.com/whonion/go-client-faucet-request

[lint-badge]: https://github.com/whonion/go-client-faucet-request/actions/workflows/lint.yml/badge.svg
[lint-url]: https://github.com/whonion/go-client-faucet-request/actions/workflows/lint.yml

[test-badge]: https://github.com/whonion/go-client-faucet-request/actions/workflows/test.yml/badge.svg
[test-url]: https://github.com/whonion/go-client-faucet-request/actions/workflows/test.yml

[build-badge]: https://github.com/whonion/go-client-faucet-request/actions/workflows/build.yml/badge.svg
[build-url]: https://github.com/whonion/go-client-faucet-request/actions/workflows/build.yml

[makefile-badge]: https://github.com/whonion/go-client-faucet-request/actions/workflows/makefile.yml/badge.svg
[makefile-url]: https://github.com/whonion/go-client-faucet-request/actions/workflows/makefile.yml

[hint-badge]: https://hits.dwyl.com/whonion//go-client-faucet-request.svg
[hint-url]: https://hits.dwyl.com/whonion/go-client-faucet-request
