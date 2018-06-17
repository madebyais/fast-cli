FAST-CLI
--------

v0.0.1

FAST-CLI is a cli for FAST serverless framework.

## GET-STARTED

1. Go to your terminal
2. Setup your GOPATH
3. Setup your GOBIN ~> `$GOPATH/bin`
4. Run `go get -u -v github.com/madebyais/fast-cli`
5. Then go to your `$GOPATH/github.com/madebyais/fast-cli`
6. And run `make`
7. Last but not least, run `fast-cli setup`

And now you are ready to go!

## COMMANDS

`help` is a command to display the help information.

```
$ fast-cli help
```

`setup` is a command to setup FAST package dependencies.

```
$ fast-cli setup
```

`create {name}` is a command to create a .go file which contain the template for FAST.

```
$ fast-cli create hello
```

`build {name}` is a command to build the .go file and export it as .so file to be used in FAST.

```
$ fast-cli build hello
```
