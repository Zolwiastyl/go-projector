# go-projector

CLI for setting variables per directory — so you can set `env` for current directory to `dev` and to directory next to it to `prod`.
And when you go between them you will just call `$(go-projector env)` instead of thinking or forgetting if this should be `dev` or `prod`.

\_Made along with ThePrimeagen while watching polyglot programmer on Frontend Masters.\_

## Building

To install dependencies run:

```sh
go mod download
```

To test run:

```sh
go test ./...
```

To get a binary simply run:

```sh
go build -o ./go-projector
```

You may copy `./go-projector` to your desired bin directory with any name you wish — for rest of `README` I will refer to it as `go-projector`.

## Use

To set a variable run:

```sh
go-projector add foo baz
```

To show a variable just run:

```sh
go-projector foo
```

To remove a variable run:

```sh
go-projector rm foo
```

To show all variables recursively go with it:

```sh
go-projector
```

Then you will get all variables from your directory, then will get all variables from directory above and until you reach you `$HOME` directory.

### Flags

There are two flags available while running this program:

- `--pwd -p` - you can specify from which directory you want to start searching for variables
- `--config -c` - you can specify from which file you want to read variables - it defaults to `.projector.json` in your `$HOME` directory
