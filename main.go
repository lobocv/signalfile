package main

import (
	"log"
	"fmt"
	"github.com/golang/protobuf/proto"
	"signalfile/signalproto"
)

func main() {

	_signals := [3]*signalproto.Signal{{1, 10, 20}, {2, 11, 21}, {3, 12, 22}}
	var signals []*signalproto.Signal = _signals[:]


	sigfile := &signalproto.SignalFile{
		SamplePoints: 500,
		SamplingIntervalPs: 100,
		Frequency_MHz: 20,
		Signals: signals,
	}

	// Serialize the structure
	data, err := proto.Marshal(sigfile)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}

	// Allocate a new structure that we can decode the binary into
	message := &signalproto.SignalFile{}

	// Decode
	err = proto.Unmarshal(data, message)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}

	// Now sigfile and message contain the same data.
	fmt.Printf("Original Message: %v\nRecieved Message: %v\n", sigfile.GetSamplePoints(), message.GetSamplePoints())
	fmt.Printf("Original Message: %v\nRecieved Message: %v\n", sigfile.GetSamplingIntervalPs(), message.GetSamplingIntervalPs())
	fmt.Printf("Original Message: %v\nRecieved Message: %v\n", sigfile.GetFrequency_MHz(), message.GetFrequency_MHz())

	fmt.Printf("Original Message: %v\nRecieved Signal: %v\n", sigfile.GetSignals(), message.GetSignals())

	// etc.
}