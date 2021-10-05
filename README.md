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