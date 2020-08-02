# [WIP]

[![Build Status](https://travis-ci.com/mgkanani/goroutines.svg?branch=master)](https://travis-ci.com/mgkanani/goroutines)
[![Coverage Status](https://coveralls.io/repos/github/mgkanani/goroutines/badge.svg?branch=master)](https://coveralls.io/github/mgkanani/goroutines?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/mgkanani/goroutines)](https://goreportcard.com/report/github.com/mgkanani/goroutines)
[![GoDoc](https://img.shields.io/badge/Godoc-reference-blue.svg)](https://godoc.org/github.com/mgkanani/goroutines)

**Problem**:
Currently, channel based approach can be used for implementing different Producer-Consumer patterns 
like single-many/many-many/many-single. In most of the cases, it can be work efficiently.
However, in few scenarios where producers and/or consumers are fast enough then underneath channel-lock
 will be a bottleneck.


