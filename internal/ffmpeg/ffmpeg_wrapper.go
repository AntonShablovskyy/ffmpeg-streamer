package ffmpeg

import (
	"fmt"
	"log"
	"os/exec"
	"strconv"
)

type FFMpeg struct {
	streamLoop int
	hlsTime    int
	inputFile  string
	outputDir  string
}

func New(streamLoop, hlsTime int, inputFile, outputDir string) *FFMpeg {
	return &FFMpeg{
		streamLoop: streamLoop,
		hlsTime:    hlsTime,
		inputFile:  inputFile,
		outputDir:  outputDir,
	}
}

func (f *FFMpeg) Mp4ToHls() error {
	cmd := exec.Command("ffmpeg",
		"-stream_loop", strconv.Itoa(f.streamLoop),
		"-i", f.inputFile,
		"-start_number", "0",
		"-hls_time", strconv.Itoa(f.hlsTime),
		"-hls_list_size", "0",
		"-f", "hls",
		fmt.Sprintf("%s/media.m3u8", f.outputDir),
	)

	log.Println("Creating HLS...")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("Failed to create HLS: %w \n %s", err, string(output))
	}

	return nil
}
