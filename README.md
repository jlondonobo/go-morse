# Go Morse

A go-based CLI for translating strings to morse code.

## Features

- Output string representation to `stdout`.
- Output sound representation to a `.wav` file via the `--output` flag.
- Supports short flags like `-s` (sound) and `-o` (output).
- Configure sound pitch `--pitch`

## Quickstart

To translate text to Morse code, run:

```console
morse 'Vamos' -s
```

This command will write the string representation to standard output and play it on your speakers.

To name the ouput file something different, use the `--file-name` (`-f`) flag. For example:

```console
morse 'Lets go, Carlos, lets go' -sf 'carlitos.wav'
```

## Roadmap

- [x] Enable saving sound to file
- [x] Make functions run in parallel
- [x] Set up short-version flags
- [x] Extend punctuation
- [x] Add a default file name for better ergonomics
- [ ] Enable editing sound qualities

  - [ ] Speed
  - [ ] Pitch
  - [ ] Volume

- [ ] Improve efficiency of steream construction / duplication.

## Configuring sounds

By default, the sound output is a 700hz sine wave. With a speed of 20 words per minute (WPM).

Morse supports configuring the pitch of sound anywhere from 300hz to 1000hz by using the `--pitch` flag or its short form `-p`.

```console
morse 'Ace' -s --pitch 500
```

Or a shorter version

```console
morse 'Ace' -sp 500
```
