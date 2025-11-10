package sound

import (
	"log"
	"os"
	"time"

	"github.com/gopxl/beep/v2"
	"github.com/gopxl/beep/v2/generators"
	"github.com/gopxl/beep/v2/speaker"
	"github.com/gopxl/beep/v2/wav"
)

const (
	WordsPerMinute = 20
)

// Todo: Extract generator

func Play(s string) {
	sr := beep.SampleRate(48000)
	speaker.Init(sr, 4800)

	sine, err := generators.SineTone(sr, 700) // hardcode freq for now
	if err != nil {

		panic(err)
	}
	silence := generators.Silence(-1)

	// based on: https://morsecode.world/international/timing/
	spd := time.Duration(60 * 1000 / (50 * WordsPerMinute)) // in milliseconds
	dit := sr.N(time.Millisecond * spd)
	dah := dit * 3

	// todo: this is not exact correspondence. Just for convenience assuming
	// 1-dit silence after every symbol.
	var m = map[string]func() beep.Streamer{
		".": func() beep.Streamer { return beep.Take(dit, sine) },
		"-": func() beep.Streamer { return beep.Take(dah, sine) },
		" ": func() beep.Streamer { return beep.Take(dit*2, silence) },
		"/": func() beep.Streamer { return beep.Take(0, silence) },
	}

	ch := make(chan struct{})
	var sounds []beep.Streamer

	for _, v := range s {
		st, ok := m[string(v)]
		if !ok {
			panic("Unrecognize symbol.")
		}
		sounds = append(sounds, st(), beep.Take(dit, silence))
	}
	sounds = append(sounds, beep.Callback(func() {
		ch <- struct{}{}
	}))
	speaker.Play(beep.Seq(sounds...))
	<-ch
	time.Sleep(200 * time.Millisecond) // to ensure last signal plays
}

func Write(s string, outfile string) {
	sr := beep.SampleRate(48000)
	speaker.Init(sr, 4800)

	sine, err := generators.SineTone(sr, 700) // hardcode freq for now
	if err != nil {

		panic(err)
	}
	silence := generators.Silence(-1)

	// based on: https://morsecode.world/international/timing/
	spd := time.Duration(60 * 1000 / (50 * WordsPerMinute)) // in milliseconds
	dit := sr.N(time.Millisecond * spd)
	dah := dit * 3

	// todo: this is not exact correspondence. Just for convenience assuming
	// 1-dit silence after every symbol.
	var m = map[string]func() beep.Streamer{
		".": func() beep.Streamer { return beep.Take(dit, sine) },
		"-": func() beep.Streamer { return beep.Take(dah, sine) },
		" ": func() beep.Streamer { return beep.Take(dit*2, silence) },
		"/": func() beep.Streamer { return beep.Take(0, silence) },
	}

	var sounds []beep.Streamer

	for _, v := range s {
		st, ok := m[string(v)]
		if !ok {
			panic("Unrecognize symbol.")
		}
		sounds = append(sounds, st(), beep.Take(dit, silence))
	}

	finalStreamer := beep.Seq(sounds...)
	f, err := os.Create(outfile)
	if err != nil {
		panic("Unable to create file.")
	}
	fmt := beep.Format{SampleRate: 48000, NumChannels: 2, Precision: 2}
	err = wav.Encode(f, finalStreamer, fmt)
	if err != nil {
		log.Fatal(err)
	}

}
