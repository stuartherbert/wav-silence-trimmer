package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"os"

	"github.com/go-audio/wav"
)

func main() {
	input := flag.String("i", "", "Input WAV file")
	output := flag.String("o", "", "Output WAV file")
	threshold := flag.Float64("t", 0.01, "Silence threshold (amplitude)")
	flag.Parse()

	if *input == "" || *output == "" {
		fmt.Printf("Usage: %s -i input.wav -o output.wav [-t threshold]\n", os.Args[0])
		return
	}

	// Read the WAV file
	f, err := os.Open(*input)
	if err != nil {
		log.Fatalf("unable to open file: %v", err)
	}
	defer f.Close()

	decoder := wav.NewDecoder(f)

	// Convert audio data to float64 samples
	pcmBuf, err := decoder.FullPCMBuffer()
	if err != nil {
		log.Fatalf("unable to read audio from file: %v", err)
	}
	f64Buf := pcmBuf.AsFloatBuffer()

	// Find first and last non-silent sample
	start, end := 0, len(f64Buf.Data)-1
	thresholdAmplitude := threshold

	// Find first non-silent sample
	for i := range f64Buf.Data {
		if math.Abs(f64Buf.Data[i]) > *thresholdAmplitude {
			start = i
			break
		}
	}

	// Find last non-silent sample
	for i := end; i >= 0; i-- {
		if math.Abs(f64Buf.Data[i]) > *thresholdAmplitude {
			end = i
			break
		}
	}

	if start > end {
		log.Printf("Entire file is silence, output will be empty")
		end = start
	}

	// Trim the samples
	pcmBuf.Data = pcmBuf.Data[start : end+1]

	fw, err := os.Create(*output)
	if err != nil {
		log.Fatalf("couldn't create output file: %v", err)
	}

	encoder := wav.NewEncoder(
		fw,
		pcmBuf.Format.SampleRate,
		int(decoder.BitDepth),
		pcmBuf.Format.NumChannels,
		int(decoder.WavAudioFormat),
	)
	err = encoder.Write(pcmBuf)

	if err != nil {
		log.Fatalf("Error writing WAV file: %v", err)
	}
	err = encoder.Close()
	if err != nil {
		log.Fatalf("Error writing WAV file: %v", err)
	}

	fw.Close()

	fmt.Printf("Trimmed %d samples\n", len(f64Buf.Data)-len(pcmBuf.Data))

	// calculate trimmed area
	trimmedBeginning := float64(start) / float64(pcmBuf.Format.SampleRate) / float64(pcmBuf.Format.NumChannels)
	trimmedEnd := (float64(len(f64Buf.Data)-1) - float64(end)) / float64(pcmBuf.Format.SampleRate) / float64(pcmBuf.Format.NumChannels)

	fmt.Printf("Trimmed  leading silence: %d samples; %0.2f sec(s)\n", start, trimmedBeginning)
	fmt.Printf("Trimmed trailing silence: %d samples; %0.2f sec(s)\n", len(f64Buf.Data)-1-end, trimmedEnd)
}
