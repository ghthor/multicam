package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	framerate := flag.Uint("framerate", 25, "cam input framerate")
	videoSizeW := flag.Uint("w", 640, "video size - Width")
	videoSizeH := flag.Uint("h", 480, "video size - Height")
	inputFormat := flag.String("input_format", "mjpeg", "cam input format")
	dev := flag.String("i", "/dev/video0", "cam device file")

	flag.Parse()

	output := filepath.Base(*dev) + ".mkv"

	ffmpeg, err := exec.LookPath("ffmpeg")
	if err != nil {
		log.Fatal(err)
	}

	//ffmpeg -f v4l2 -framerate 25 -video_size 640x480 -input_format mjpeg -i /dev/video0 video0.mkv
	cmd := exec.Command(ffmpeg,
		"-f", "v4l2",
		"-framerate", fmt.Sprint(*framerate),
		"-video_size", fmt.Sprintf("%vx%v", *videoSizeW, *videoSizeH),
		"-input_format", *inputFormat,
		"-i", *dev,
		output)

	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	fmt.Println(cmd)

	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
