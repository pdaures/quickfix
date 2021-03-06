QuickFIX/Go [![GoDoc](https://godoc.org/github.com/quickfixgo/quickfix?status.png)](https://godoc.org/github.com/quickfixgo/quickfix) [![Build Status](https://travis-ci.org/quickfixgo/quickfix.svg?branch=master)](https://travis-ci.org/quickfixgo/quickfix) [![Go Report Card](https://goreportcard.com/badge/github.com/quickfixgo/quickfix)](https://goreportcard.com/report/github.com/quickfixgo/quickfix)
===========

Open Source [FIX Protocol](http://www.fixprotocol.org/) library implemented in Go

FIX versions 4.0-5.0

Example Apps
------------

See [examples](https://github.com/quickfixgo/examples) for some simple examples of using QuickFIX/Go.

Build and Test
--------------

The default make target runs [go vet](https://godoc.org/golang.org/x/tools/cmd/vet) and unit tests.

QuickFIX/Go acceptance tests depend on ruby in path.

To run acceptance tests,

		# build acceptance test rig
		make build_accept

		# run acceptance tests
		make accept
