package main

import (
	"log"
	"fmt"
	"github.com/golang/protobuf/proto"
	"signalfile/signalproto"
)


type SignalFile struct {

	signal *signalproto.Signal

}

func NewSignalFile(signal *signalproto.Signal) SignalFile  {
	sigfile := SignalFile{signal}
	return sigfile
}


func (sigfile SignalFile) Serialize() []byte {
	// Serialize the structure into a byte array
	data, err := proto.Marshal(sigfile.signal)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	return data
}

func Deserialize(data []byte) signalproto.Signal {
	// Return a Signal structure from it's byte representation
	signal := signalproto.Signal{}
	err := proto.Unmarshal(data, &signal)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}
	return signal
}


func main() {

	_signals := [3]*signalproto.SignalData{{1, 10, 20}, {2, 11, 21}, {3, 12, 22}}
	var signals []*signalproto.SignalData = _signals[:]


	signal := &signalproto.Signal{
		SamplePoints: 500,
		SamplingIntervalPs: 100,
		Frequency_MHz: 20,
		Signals: signals,
	}

	sigfile := NewSignalFile(signal)
	data := sigfile.Serialize()


	message := Deserialize(data)

	// Now signal and message contain the same data.
	fmt.Printf("Original Message: %v\nRecieved Message: %v\n", signal.GetSamplePoints(), message.GetSamplePoints())
	fmt.Printf("Original Message: %v\nRecieved Message: %v\n", signal.GetSamplingIntervalPs(), message.GetSamplingIntervalPs())
	fmt.Printf("Original Message: %v\nRecieved Message: %v\n", signal.GetFrequency_MHz(), message.GetFrequency_MHz())

	fmt.Printf("Original Message: %v\nRecieved Signal: %v\n", signal.GetSignals(), message.GetSignals())

	// etc.
}