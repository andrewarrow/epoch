package sound

/*
import (
	"os"

	"github.com/hajimehoshi/oto"
	"github.com/tosone/minimp3"
)

func PlaySound(filename string) {
	if true {
		return
	}
	asBytes, _ := os.ReadFile("assets/sounds/done5.mp3")
	decoder, data, _ := minimp3.DecodeFull(asBytes)
	context, _ := oto.NewContext(decoder.SampleRate, 2, 2, 4096)

	defer context.Close()
	player := context.NewPlayer()
	defer player.Close()
	player.Write(data)
} */
