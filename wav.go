package wavPlayer

import (
    "errors"
	"fmt"
	"io"
	"os"
	"github.com/cocoonlife/goalsa"
	"github.com/cryptix/wav"
)

type WavPlayer struct {
    stream *[]int16
    device *alsa.PlaybackDevice
}

func NewWAV(filename string) (*WavPlayer, error) {
    player := WavPlayer{}
    soundfile, err := os.Open(filename)
    if err != nil {
        return nil, err
    }

    // stat for size
    sndfileinfo, err := os.Stat(soundfile.Name())

    if err != nil {
        return nil, errors.New(fmt.Sprint("stat:", err))
    }

    // wavReader
    wavReader, err := wav.NewReader(soundfile, sndfileinfo.Size())
    if err != nil {
        return nil, errors.New(fmt.Sprint("WAV reader:", err))
    }

    // require wavReader
    if wavReader == nil {
        return nil, errors.New(fmt.Sprint("nil wav reader"))
    }

    fileinfo := wavReader.GetFile()

    samplerate := int(fileinfo.SampleRate)
    if samplerate == 0 || samplerate > 100000 {
        samplerate = 44100
    }

    player.device, err = alsa.NewPlaybackDevice("default", 1, alsa.FormatS16LE, samplerate, alsa.BufferParams{})

    if err != nil {
        return nil, errors.New(fmt.Sprint("alsa:", err))
    }

    // Require ALSA device
    if player.device == nil {
        return nil, errors.New("Nil Alsa device")
    }

    for {
        s, err := wavReader.ReadSampleEvery(2, 0)
        var cvert []int16
        for _, b := range s {
            cvert = append(cvert, int16(b))
        }
        player.stream = &cvert

        if err == io.EOF {
            break
        } else if err != nil {
            return nil, errors.New(fmt.Sprint("WAV Decode:", err))
        }
    }

    return &player, nil
}

func (player *WavPlayer) Play() {
    if player.device == nil {
        return
    }
    if *player.stream != nil {
        player.device.Write(*player.stream)
    }
}

func (player *WavPlayer) Close() {
    player.device.Close()
}
