# WAV Silence Trimmer

## Introduction

This is a quick-and-dirty CLI tool. Use it to trim leading and trailing silence from your WAV files.

I've built this to automate one of the steps I take when publishing audio demos for [the HomeToneBlog](http://hometoneblog.com).

## How To Build

Follow these steps:

1. clone the repository
2. run `go build`

You should end up with a `trimmer` CLI tool in the top-level folder.

## How To Use

`./trimmer -i <input-file> -o <output-file> -t <silence threshold>`

where:

- `<input-file>` is the WAV file that you want to trim
- `<output-file>` is the new file that `trimmer` will write out
- `<silence threshold>` is the per-sample value for audio silence

With my own audio demos, I've found `-t 950` (or thereabouts) is a good threshold for detecting 'silence'. Experiment to find an effective threshold for your recordings.

(Turns out, if you're processing music exported from a DAW, that 'silent' leading audio isn't actually empty. The samples are just too low a value to be heard!).

## Future Plans

At some point (when it starts to annoy me - so probably soon!), I'll update WAV Silence Trimmer to support processing batches of files.

## Contributions Welcome

If you've got an idea for how to improve this tool, hit me up. I'm not looking to expand the scope of the tool, but I'm definitely open to anything that will make the tool do a better job of trimming leading and trailing silence.