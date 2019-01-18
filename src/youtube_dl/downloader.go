package youtube_dl

import (
	"fmt"
	// "github.com/BrianAllred/goydl"
	"github.com/rylio/ytdl"
	// "io"
	// "log"
	"os"
)

func Downloader(videoUrl string) {
	vid, err := ytdl.GetVideoInfo("https://www.youtube.com/watch?v=" + videoUrl)
	fmt.Println(*vid, "Downloader executed")
	fmt.Println((*vid).Formats, "Formats")
	formatsList := (*vid).Formats

	bestFormatsList := formatsList.Best("res")
	fmt.Println(bestFormatsList, "best Formats")

	if err != nil {
		fmt.Println(err, "Error")
	}
	file, _ := os.Create(vid.Title + ".mp4")
	defer file.Close()
	downloadError := vid.Download(bestFormatsList[0], file)

	if downloadError != nil {
		fmt.Println(downloadError, "downloadError")
	}
}

/*func Downloader(videoUrl string) {
	fmt.Println(videoUrl, "Downloader executed")
	youtubeDl := goydl.NewYoutubeDl()

	// youtubeDl.Options.Output.Value = "./sp.mp3"
	// youtubeDl.Options.ExtractAudio.Value = true
	// youtubeDl.Options.AudioFormat.Value = "mp3"

	// // Or update the binary
	// youtubeDl.Options.Update.Value = true

	// // Optional, required if binary is not in $PATH
	// youtubeDl.YoutubeDlPath = "/home/ravindra/projects/go-projects/youtube-downloader"

	// go io.Copy(os.Stdout, youtubeDl.Stdout)
	// go io.Copy(os.Stderr, youtubeDl.Stderr)

	cmd, err := youtubeDl.Download("https://www.youtube.com/watch?v=" + videoUrl)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Title: %s\n", youtubeDl.Info.Title)
	cmd.Wait()
}*/

// WriteToFile(w io.Writer) {
//     b, _ := json.Marshal(*p)
//     w.Write(b)
// }
