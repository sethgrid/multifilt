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

Given the files the above example (`In` and `Filter`)
```
$ $ cat In | mf Filter | diff Out - # Out is the file provided in the cmd/mf dir
$ echo $? # 0; meaning that Out is as expected
$ # typical usages
$ cat In | mf Filter > Out
$ # or
$ mf -in In -f Filter -out Out
```