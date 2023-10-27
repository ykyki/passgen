# passgen

A simple password generator.

## Alternative

shell script
```zsh
function passgen {
    local passlen=${1:-32}
    LC_ALL=C tr -cd '[:alnum:]!@#$%^&*' < /dev/urandom | fold -w $passlen | head -n1
}
```

## Usage

```sh
$ passgen -h
Usage of passgen:
  -v, --version      print passgen version
  -l, --length int   password length (default 32)
  -A, --capital      include capital letters
  -a, --small        include small letters
  -n, --number       include numbers
  -s, --symbol       include symbols

$ passgen
&9M8F5Sui1FzFqZ@LfNbgBHOz2zynd%o

$ passgen -aA -l 16
aiOYZsoCoeirOBtQ
```

## Installation

```sh
go install github.com/ykyki/passgen@latest
```
