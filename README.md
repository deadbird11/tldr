# tldr
## Usage
`tldr "command"`
## About
Command line tool that gives a concise explanation of a given command, along with example usage. I got this idea from [this](https://www.reddit.com/r/golang/comments/7pnw2e/fun_golang_projects/) reddit post, and is pretty much just a copy of [github.com/mstruebing/tldr](https://github.com/mstruebing/tldr). Copying this is not very creative, but I just wanted an excuse to get started in Go and I didn't want to wait for "inspiration." That being said, this is entirely my own code.
## Get Started
### Windows
Clone this repository. Then follow the steps on [this tutorial](https://medium.com/@kevinmarkvi/how-to-add-executables-to-your-path-in-windows-5ffa4ce61a53), to add `tldr.exe` to your PATH. You should now be able to use `tldr` on the command line!
### Otherwise
You will need to have a recent version of Go installed on your machine. Build the executable with `go build tldr.go`, then add the `tldr.exe` file to your PATH.
