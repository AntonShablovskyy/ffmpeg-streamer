package main

import (
	"errors"
	"ffmpeg-streamer/internal/ffmpeg"
	"ffmpeg-streamer/internal/server"
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	port       int
	streamLoop int
	hlsTime    int
	inputFile  string
	outputDir  string
)

func main() {
	// read and validate command-line arguments
	err := parseArgs()
	if err != nil {
		log.Fatal(fmt.Sprintf("Error parsing command-line arguments: %s", err))
	}

	// create hls output directory if not exists
	err = createOutputDir()
	if err != nil {
		log.Fatal(fmt.Sprintf("Can't create HLS output directory: %s", err))
	}

	// create FFMPEG client wrapper and generate HLS
	f := ffmpeg.New(streamLoop, hlsTime, inputFile, outputDir)
	err = f.Mp4ToHls()
	if err != nil {
		log.Fatal(fmt.Sprintf("Error creating HLS: %s", err))
	}

	// serve HLS via HTTP
	s := server.New(port, outputDir)
	s.ServeHls()
}

func parseArgs() error {
	flag.IntVar(&port, "port", 8080, "HTTP expose port. Default value is 8080")
	flag.IntVar(&streamLoop, "streamLoop", 0, "Set number of times input stream shall be looped. Default value is 0")
	flag.IntVar(&hlsTime, "hlsTime", 2, "Set the target segment length. Default value is 2")
	flag.StringVar(&inputFile, "inputFile", "", "Input mp4 file to stream")
	flag.StringVar(&outputDir, "outputDir", "hls_output", "Output directory for storing HLS file")

	flag.Parse()

	if inputFile == "" {
		return errors.New("missing required argument 'inputFile'")
	}

	return nil
}

func createOutputDir() error {
	if _, err := os.Stat(outputDir); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(outputDir, os.ModePerm)
		if err != nil {
			return err
		}
	}

	return nil
}
