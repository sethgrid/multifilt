# MultiFilt

`multifilt` allows you to filter out multiple lines of a file (or any `io.Reader`) based on multiple lines of another file.

## Example
`multifilt` allows you to filter on partial line matches or full line matches.
Here is what would happen with a partial line match.
The `In` and `Filter` columns represent each an `io.Reader`, and `Out` is an `io.Writer`.
```
In:           Filter:   Out:
The cat       at        with a plan
in the hat    met       Panama!
met a man
with a plan
Panama!
```

# Installing `mf`, the commandline tool

The `mf` commandline tool can be installed with:
`go get -u github.com/sethgrid/multifilt/...`

## usage

```
$ mf -h
Usage: mf (multifilter)

Version 1.2.0, Compiled with go1.7.1

Description
Filter out lines from an input source based on lines in a filter file and/or -v flags.
A filter file with two lines, 'ab' and 'bc', will filter out lines from the input source that contain either entry.

Examples
cat input | mf filter_file -a > output
cat input | mf -v foo -v bar -v raz filter_file > output
mf -in input -out ouput -f filter_file

  -a	filtered lines must match the whole line in the filter ('a' for match all)
  -f string
    	file filter, use -f or provide as single argument
  -in string
    	file in, default stdin
  -out string
    	file out, default stdout
  -v value
    	specify multiple -v params to filter on each
```