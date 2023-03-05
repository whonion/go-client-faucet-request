 [![Build Status](https://travis-ci.org/joemccann/dillinger.svg?branch=master)](https://travis-ci.org/joemccann/dillinger) [![Go version][go-badge]][go-url] [![HitCount](https://hits.dwyl.com/whonion//go-client-faucet-request.svg)](https://hits.dwyl.com/whonion/go-client-faucet-request)
# Faucet-client for send multi-request to server using proxy</br>
Implementation with go-routines
## Package description
 - main.go - main package to run or build
 - proxy.go - package for getting a proxy list from a remote site
 - proxy.txt - proxy list file in the protocol://ip:port@login@pass format
 - batch generator of wallets with output of 'addresses.txt' file
 - addresses.txt - wallet's list for sending curl requests to target faucet's url
 - agents.txt - list of user-agents


[go-badge]: https://img.shields.io/badge/go-1.20-blue.svg
[go-url]: https://go.dev
