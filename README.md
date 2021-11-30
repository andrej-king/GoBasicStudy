### First GO project

<p>
    :information_source: <b>GOROOT</b> and <b>GOPATH</b> should already be configured.
</p>

<p>
To start a new project in Go, first need create an empty folder.<br>
Then write the <b>go modules init</b> command in this folder to initialize the Go modules.<br>
Modules are used for dependencies.
</p>

`mkdir hello-world` \
`cd hello-world` \
`go mod init github.com/[username]/[project name]`

`go mod tidy` Download missing packages

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

### Multithreading
`go` Before calling the method, will run it on a new thread <font color="red">(If the main thread has finished executing, all other threads will not be executed.)</font> \
`channels` use for multithreading \
`select` Allow read data from multiple channels. \
`sync.WaitGroup` Allow waiting end all goroutines in one moment (together). \
`sync\atomic` Thread safe counter.

### Books for more information
[golang-book.com](https://golang-book.com) original (eng) \
[golang-book.ru](https://golang-book.ru) translated (rus)
