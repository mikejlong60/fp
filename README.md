# Using Code Generator

```
go clean
go generate
go test
```

You can add more types in the special compiler directive at the top of concrete_either.go

But you will get some duplicated structures.  For now delete them by hand from the generated code. We will deduplicate the code generator later.

