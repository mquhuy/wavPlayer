# WAV Player for Golang

This repo is a replicate of https://github.com/aerth/playwav

The `playwav` repo is a great way to play wav files from command line, or to make music apps where each audio file is run only once, but for purposes like games, where every sound file can be played several times, it's too slow, since every time a file is played, it gets loaded from the disk.

This repo serves uses the same logic to create a wav player as the original one, but adding adaptation that allows replaying the audio without the need of reloading the file, hence suitable for purposes like game-making.

## How to use
- Get the library: `go get "github.com/lthh91/WAVPlayer"`
- Add `github.com/lthh91/WAVPlayer` to the import part in your `GO` file
- Create the `Player` object with `player := wavAudio.NewWAV()`
- Play the audio with `player.Play()`
- Kill the player using `player.Close()`

*Example*: https://github.com/lthh91/FlappyBird/blob/master/bird.go#L31
