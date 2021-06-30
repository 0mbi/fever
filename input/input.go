package input

// DCSO FEVER
// Copyright (c) 2017, DCSO GmbH

// Input describes how to handle
// events parsed from EVE input.
type Input interface {
	GetName() string
	Run()
	SetVerbose(bool)
	Stop(chan bool)
}
