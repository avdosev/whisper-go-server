package main

import (
	"fmt"
	"io"
	wav "github.com/go-audio/wav"
	// "github.com/ggerganov/whisper.cpp/bindings/go/pkg/whisper"
	"github.com/avdosev/whisper.cpp/bindings/go/pkg/whisper"

)

var m whisper.Model

func init_model() {
	var modelpath string = "../whisper.cpp/models/ggml-medium-q4_1.bin" 
	modelpath = "../whisper.cpp/models/ggml-small-q4_1.bin" // TODO: use medium m
	var err error

	// load m
	m, err = whisper.New(modelpath)
	if err != nil {
		panic(err)
	}
}


func audio_transcribe(file io.ReadSeeker) (string, error) {
	var samples []float32 // Samples to process

	// Process samples
	context, err := m.NewContext()

	if err != nil {
		return "", err
	}

	context.SetLanguage("ru")
	context.SetTranslate(false)

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