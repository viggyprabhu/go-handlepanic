go-handlepanic
===========

A Golang library for handling various panic in one central location.

See [example](https://github.com/viggyprabhu/panicSample/blob/master/src/panicSample/main.go) for usage.

We wrap our go routine call inside `handlepanic.Handle` instead of calling it directly as shown below:

We use the following 

```go
args := Args{
    a: 1,
    b: 0,
}
handlepanic.CreateInstance(rescue)
go handlepanic.Handle(panicRoutine, args)
```

instead of 

```go
go panicRoutine(args)
```

where `rescue` is a function which handles panic in a single location. 
