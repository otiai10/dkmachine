# dkmachine

[![Build Status](https://travis-ci.org/otiai10/dkmachine.svg?branch=master)](https://travis-ci.org/otiai10/dkmachine)

Go SDK API accessor for `docker-machine`.

```go
opt := &dkmachine.CreateOptions{
  Name:   "foobar",
  Driver: "virtualbox",
}

machine, err := dkmachine.Create(opt)
```

# Test

Testing requires VirtualBox on your environment.