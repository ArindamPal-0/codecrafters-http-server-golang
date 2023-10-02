# Build your own HTTP server

[![progress-banner](https://backend.codecrafters.io/progress/http-server/28eed93f-b787-49be-a1fd-6d7f4cc90241)](https://app.codecrafters.io/users/ArindamPal-0?r=2qF)

[CodeCrafters](https://app.codecrafters.io/)

## Setup

To build the project
```shell
go build app/server.go
```

Run the project
```shell
./server --directory <directory_for_files>
```

### Dev Setup

To run the project
```shell
go run app/server.go --directory <directory_for_files>
```

## CodeCrafters Instructions

This is a starting point for Go solutions to the
["Build Your Own HTTP server" Challenge](https://app.codecrafters.io/courses/http-server/overview).

[HTTP](https://en.wikipedia.org/wiki/Hypertext_Transfer_Protocol) is the
protocol that powers the web. In this challenge, you'll build a HTTP/1.1 server
that is capable of serving multiple clients.

Along the way you'll learn about TCP servers,
[HTTP request syntax](https://www.w3.org/Protocols/rfc2616/rfc2616-sec5.html),
and more.

**Note**: If you're viewing this repo on GitHub, head over to
[codecrafters.io](https://codecrafters.io) to try the challenge.

## Passing the first stage

The entry point for your HTTP server implementation is in `app/server.go`. Study
and uncomment the relevant code, and push your changes to pass the first stage:

```sh
git add .
git commit -m "pass 1st stage" # any msg
git push origin master
```

Time to move on to the next stage!

## Stage 2 & beyond

Note: This section is for stages 2 and beyond.

1. Ensure you have `go (1.19)` installed locally
1. Run `./your_server.sh` to run your program, which is implemented in
   `app/server.go`.
1. Commit your changes and run `git push origin master` to submit your solution
   to CodeCrafters. Test output will be streamed to your terminal.
