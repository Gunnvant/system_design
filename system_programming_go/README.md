## System Programming Go

These notes are based on System Programming Essentials with Go: Alex Rios

### [Concurrency and Parallelism: Go routines](./go_routines/)

This section has notes on basic concurrency building blocks of concurrency in go are:

1. [Go routines](./go_routines/intro/): These are equivalent to threads in other languages. `WaitGroup` help in blocking the main thread till all other threads complete.

2. [Mutexes](./go_routines/mutexes/): When more than on go routine accesses a resource (memory,file etc), we need to make sure that there is no race condition. 

One can discover if there is any race condition in the code by doing unit testing and passing the `-race` flag.

```bash
go test -race
```

Once a race condition is discovered it can be tackled by:
1. Using atomic [operations](./go_routines/mutexes/atomic.go)
2. Using mutex [locks](./go_routines/mutexes/mutex_lock.go)


3. [Channels](./go_routines/channels/)

Channels provide a mechanism for communicating between go routines. There are unbuffered channels and buffered channels. 

Unbuffered channels don't have any pre-defined size.

One can create them as given in [this example](./go_routines/channels/unbuffered.go).
WaitGroups help in blocking the main goroutine while [channels are closed and listeners on the channels finish doing their tasks](./go_routines/channels/unbuffered_waitgroup.go).

In case one needs to listen to many channels one can make use of [range](./go_routines/channels/unbuffered_range.go)