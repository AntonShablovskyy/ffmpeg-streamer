# ffmpeg-streamer

ffmpeg-streamer converts input MP4 file to HLS files and expose it via HTTP.

## Dependencies
[ffmpeg executable](https://www.ffmpeg.org/download.html) must be installed on host machine.

## Build
`go build .`

## Command-line arguments
+ **inputFile** : Input mp4 file to stream. **Required argument.**
+ **port** : HTTP expose port. Default value is 8080.
+ **streamLoop** : Set number of times input stream shall be looped. Default value is 0.
+ **hlsTime** : Set the target segment length. Default value is 2.
+ **outputDir** : Output directory for storing HLS file.

## Example run
`./ffmpeg-streamer --inputFile video.mp4 --streamLoop 2 --hlsTime 10`