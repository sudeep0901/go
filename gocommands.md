go env
go help gopath        

# Build with different name
go build -o hello_world hello.go
go help modules


## goimports for 

There’s an enhanced version of go fmt available called goimports that also cleans up your import statements. It puts them in alphabetical order, removes unused imports, and attempts to guess any unspecified imports. Its guesses are sometimes inaccurate, so you should insert imports yourself.

You can download goimports with the command go install golang.org/x/tools/cmd/goimports@latest. You run it across your project with the command:

```sh
go install golang.org/x/tools/cmd/goimports@latest
goimports -l -w .


go env GOROOT
```

