# fly
[![Build Status](https://travis-ci.org/cizixs/fly.svg?branch=master)](https://travis-ci.org/cizixs/fly)
[![Coverage Status](https://coveralls.io/repos/github/cizixs/fly/badge.svg?branch=master)](https://coveralls.io/github/cizixs/fly?branch=master)
[![GoDoc](https://godoc.org/github.com/cizixs/fly?status.svg)](https://godoc.org/github.com/cizixs/fly)
[![Go Report Card](https://goreportcard.com/badge/github.com/cizixs/fly)](https://goreportcard.com/report/github.com/cizixs/fly)

A time and date toolbox for golang.

golang provides `time` library to deal with time, and it works well. 
`fly` aims to provide some convenient features for real-life time process.


## Features

- [x] humanize date/time output
- [ ] fully compatible with standard `time` library
- [ ] easy conversion between different timezones
- [ ] parse date/time in any user-defined format string
- [ ] abundant shortcuts for real-life date conversion, such as `Tomorrow`, `Sunday`, `NextYear`...
- [ ] access date ranges, floors and ceilings in all time frames(microsecond up to year)

## Usage

### Quick utc time and local time instance

When getting `now` time, users can easily access to local time and UTC time:

```
l := fly.Now()    // Now time with current location timezone if accessible
u := fly.UTCNow() // UTC timezone now
```

### Compatible with standard `time.Time`

If you have a standard `time.Time` instance, it can be converted to `fly` struct:

```
now := time.Now()
f := fly.New(now)
```

### Add duration the way you like

Need to travel in time(move time forward and backward)? `Add` method can help:

```
f := fly.Now()
f, err := f.Add(time.Duration(2*time.Hour))    // set time to 2 hours later
f, err := f.Add(time.Duration(30 * 24 * time.Hour)) // how about 30 days later
f, err := f.Add(time.Duration(100 * 12 * 30 * 24 * time.Hour)) // Around 100 years later, this is time travelling!
```

This is nothing, afterall calculating time and date is boring. See the magic:

```
f := fly.Now()
f, err := f.Add("2h")       // 2 hours later
f, err := f.Add("-2h30m")   // 2 and half hours in the past
f, err := f.Add("300s")    // 300 seconds later
```

Right now, `fly` supports `ns`, `us`, `ms`, `s`, `m`, `h` and their combinations.
If the given parameter can not be parsed as time duration, error will be retured.

### Humanize output

If you need to human readble time output, `Humanize` method is here to help:

```
f := f.Now()
f.Add("2h2m")
f.Humanize()
// 2 hours from now
```

## License

[MIT License](https://github.com/cizixs/fly/blob/master/LICENSE)
