# Go Morse

A go-based CLI for translating strings to morse code.

## Features

- Output string representation to `stdout`.
- Output sound representation to a `.wav` file via the `--output` flag.
- Supports short flags like `-s` (sound) and `-o` (output).

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
