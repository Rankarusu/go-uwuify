<div align="center">
<img height="50" src="https://github.com/egonelbre/gophers/blob/master/icon/emoji/gopher-blushing.png?raw=true">
<h1>go-uwuify</h1>
</div>

A command line tool written in Go with [Cobra](https://github.com/spf13/cobra) to make text more uwu

## Installation

```sh
go install github.com/rankarusu/go-uwuify@latest
```

This installs to $HOME/go/bin/ by default, which can be added to your $PATH:

```sh
export PATH=$PATH:$HOME/go/bin
```

## Usage

transforms the given input and outputs it to the desired destination
e.g. `uwuify --infile input.txt -o ~/output.txt -r 1`

input can be defined via `--infile`, `--text`, or stdin via the `|` operator
the transformed text will be output to stdout by default or written to a file by using `--outfile`

the modifiers `--uwu`, `--kaomoji`, `--stutters`, `--exclamations`, and `--actions` can be used to set the probability for the corresponding transforms
0 -> will not occur
.5 -> will occur 50% of the time
1 -> will occur at every possibility, usually for, or after every word

```
uwuify [flags]
```

### Options

```
  -t, --text string          a text to uwuify
  -i, --infile string        a file to uwuify
  -o, --outfile string       a file to output the uwuified text to
  -u, --unicode              allow unicode characters for kaomoji
  -r, --replacements float   probability for transforming text. e.g. love -> wuv (default 0.5)
  -k, --kaomoji float        probability for inserting kaomoji. e.g. OwO (default 0.025)
  -s, --stutters float       probability for adding stutters to the beginning of a word. e.g. hello -> h-hello (default 0.025)
  -e, --exclamations float   probability for transforming punctuation. e.g. 1 -> !!11 (default 0.5)
  -a, --actions float        probability for adding actions. e.g. *blushes* (default 0.025)
  -h, --help                 help for uwuify
```
