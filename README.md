### First GO project

<p>
    :information_source: <b>GOROOT</b> and <b>GOPATH</b> should already be configured.
</p>

<p>
To start a new project in Go, first need create an empty folder.<br>
Then write the <b>go modules init</b> command in this folder to initialize the Go modules.<br>
Modules are used for dependencies.
</p>

<pre>
mkdir hello-world
cd hello-world
go mod init github.com/[username]/[project name]
</pre>

Download missing packages:
<pre>
go mod tidy
</pre>

### Envs
`GOOS=linux;GOARCH=amd64` env for linux\
`GOOS=windows;GOARCH=amd64` env for windows\
`GOOS=darwin;GOARCH=amd64` env for macos

### Build
`extract GOROOT=D:\Go\go1.17.1 #gosetup`\
`extract GOPATH=D:\Go\go1.17.1\bin #gosetup`\
`D:\Go\go1.17.1\bin\go.exe build -o D:\GoProjects\syntax\build\lesson_1\app_windows.exe D:\GoProjects\syntax\lesson_1\app.go #gosetup`

### Tests
`cd` In folder with tests and write in terminal: \
`go test -v` Display tests result \
`go test -json` Display tests result in JSON format \
`go test -v -run /simple` Run only test with name "simple" \
`go test -v -run TestMultiple/simple` Run only method "TestMultiple with test name "simple" and all nested tests \
`go test -v -run /simple/1` Running all nested methods with this name

Add `t.Parallel()` in each test, to run tests in parallel