protoc-gen-error
================
It is a bug reporting repository for issue in [protostuff/protobuf-jetbrains-plugin](https://github.com/protostuff/protobuf-jetbrains-plugin)

Issue link: [#]() 

Details
-------

**Describe the bug**

The plugin is not able to resolve the options when used with curly braces `{`

**To Reproduce**
Check the [`types.proto`](./types.proto), it defines the options which are used in [`check.proto`](./check.proto).

You can see the Unresolved reference error in `errored` named field in `Object` message.

And to check the correctness of the syntax used in `check.proto`, just run
```shell script
go generate
```  
