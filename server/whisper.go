package main

import (
	"fmt"
	"io"
	wav "github.com/go-audio/wav"
	"github.com/ggerganov/whisper.cpp/bindings/go/pkg/whisper"
)

var model whisper.Model

func init_model() {
	var modelpath string = "../whisper.cpp/models/ggml-medium-q4_1.bin" 
	modelpath = "../whisper.cpp/models/ggml-small-q8_0.bin" // TODO: use medium model
	var err error

	// load model
	model, err = whisper.New(modelpath)
	if err != nil {
		panic(err)
	}
}

func audio_transcribe(file io.ReadSeeker) (string, error) {
	var samples []float32 // Samples to process

	// Process samples
	context, err := model.NewContext()

	if err != nil {
		return "", err
	}

	// Decode the WAV file - load the full buffer
	dec := wav.NewDecoder(file)
	if buf, err := dec.FullPCMBuffer(); err != nil {
		return "", err
	} else if dec.SampleRate != whisper.SampleRate {
		return "", fmt.Errorf("unsupported sample rate: %d", dec.SampleRate)
	} else if dec.NumChans != 1 {
		return "", fmt.Errorf("unsupported number of channels: %d", dec.NumChans)
	} else {
		samples = buf.AsFloat32Buffer().Data
	}


	if err := context.Process(samples, nil, nil); err != nil {
		return "", err
	}

	// Print out the results
	var result string = "" 
	for {
		segment, err := context.NextSegment()
		if err != nil {
			break
		}
		fmt.Printf("[%6s->%6s] %s\n", segment.Start, segment.End, segment.Text)
		result += segment.Text
	}

	return result, nil
}