protoc-gen-error
================
It is a bug reporting repository for issue in [protostuff/protobuf-jetbrains-plugin](https://github.com/protostuff/protobuf-jetbrains-plugin)

Issue link: [#155](https://github.com/protostuff/protobuf-jetbrains-plugin/issues/155) 

Details
-------
**Describe the bug**
The plugin is not able to resolve the options when used with curly braces `{`

**To Reproduce**
The complete bug is reproduced in [Shivam010/protoc-gen-error](https://github.com/Shivam010/protoc-gen-error).

Check the [`types.proto`](https://github.com/Shivam010/protoc-gen-error/blob/master/types.proto), it defines the options which are used in [`check.proto`](https://github.com/Shivam010/protoc-gen-error/blob/master/check.proto).

You can see the Unresolved reference error in `errored` named field in `Object` message.

And to check the correctness of the syntax used in `check.proto`, just run `go generate`  

**Expected behavior**
In [`check.proto`](https://github.com/Shivam010/protoc-gen-error/blob/master/check.proto), the `errored` named field, there's an error reported on `two` and also has no suggestion for any of the fields in braces.
`string errored = 1 [(main.rule).type={one: "one", two:"two"}];` 

The same works with `working` named field, in which `{` is not used.
`string working = 2 [(main.rule).type.one="one"];`

**Screenshots**
![image](https://user-images.githubusercontent.com/29069530/91180528-727f0080-e705-11ea-8d1b-6a1bf3e3f2fd.png)

**Plugin (please complete the following information):**
 - OS: macOS Catalina Version 10.15.6
 - Plugin version: 0.13.0
 - IDE: GoLand 2020.1.4
