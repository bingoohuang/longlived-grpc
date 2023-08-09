# gRPC proto files Best Practices

from [blog](https://sonny-alvesdias.medium.com/grpc-proto-files-best-practices-2d1c6169219c)

This article is a mash-up of different sources cited in references and Pixelmatic guidelines aimed to serve as coding
and style guides for proto files in the context of gRPC services definitions.

## Coding and style guides

In short, the philosophy of the guides is:

- Be consistent
- Reuse the well-known types from Google
- But keep separation as needed to avoid breaking changes

Standard file formatting

- Keep the line length to 80 characters.
- Prefer the use of double quotes for strings.

File structure

Files should be `named lower_snake_case.proto`

All files should be ordered in the following manner:

1. License header (if applicable)
2. File overview
3. Syntax
4. Package
5. Imports (sorted)
6. File options
7. Service (1 service per file)
8. Messages (sorted in the same order of the RPCs and grouped under a comment// region $RPCNAME)

```proto
service Foo {
  rpc Bar(BarRequest) returns (BarResponse);
}

// region Bar
  message BarRequest {
    // ...
  }

  message BarResponse {
    // ...
  }
```

Note: // region comments are collapsible in Visual Code.

## Packages

Package names should be in lowercase. Package names should have unique names based on the project name and possibly
based on the path of the file containing the protocol buffer type definitions.

## Message and field names

Use `CamelCase` (with an initial capital) for message names — for example, SongServerRequest. Use
underscore_separated_names for field names (including oneof field and extension names) – for example, `song_name`.

```proto
message SongServerRequest {
  optional string song_name = 1;
}
```

Using this naming convention for field names gives you accessors like the following:

C++:

```c++
  const string& song_name() { ... }
  void set_song_name(const string& x) { ... }
```

Java:

```java
  public String getSongName() { ... }
  public Builder setSongName(String v) { ... }
```

If your field name contains a number, the number should appear after the letter instead of after the underscore. For
example, use `song_name1` instead of `song_name_1`

## Optional fields

Consider the following proto3 message which defines the field bar:

```proto
*message* Foo {
  int32 bar = 1;
}
```

With this definition, it is impossible to check whether `bar` has been set to 0 or if no value has been set since the
default value of `int32` fields is `0`.

To allow you to make the difference, use:

```proto
message Foo {
  optional int32 bar = 1;
}
```

This exposes `hasBar()` and `clearBar()` (depending on the language) methods in the generated code.

## Repeated fields

Use pluralized names for repeated fields.

```proto
repeated string keys = 1;
  ...
  repeated MyMessage accounts = 17;
```

## Oneof

Oneof is a wonderful example where a protobuf language feature helps to make gRPC APIs more intuitive. As an example,
imagine we have a service method where users are able to change their profile picture, either from an URL or by
uploading their own (small) image. Instead of doing this

```proto
// Either set image_url or image_data. Setting both will result in an error.
message ChangeProfilePictureRequest {
  string image_url = 1;
  bytes image_data = 2;
}
```

we can define the desired behavior directly into the message with oneof

```proto
message ChangeProfilePictureRequest {
  oneof image {
    string url = 1;
    bytes data = 2;
  }
}
```

Not only is that much clearer for API consumers, but it is also easier to check which field has been set in the
generated code. Keep in mind that oneof also allows that none of the fields has been set, meaning there is no need to
introduce a separate none field if the oneof should be optional.

## Flexible data

If you do not know in advance the nature of the data you are going to receive from your consumer, we suggest you use the
well know type Struct. For example:

```proto
message StructTest {
  google.protobuf.Struct data = 1;
}
```

## Enums

Use `CamelCase` (with an initial capital) for enum type names and `CAPITALS_WITH_UNDERSCORES` for value names:

```proto
enum FooBar {
  FOO_BAR_UNSPECIFIED = 0;
  FOO_BAR_FIRST_VALUE = 1;
  FOO_BAR_SECOND_VALUE = 2;
}
```

- Each enum value should end with a semicolon, not a comma. Prefer prefixing enum values instead of surrounding them in
  an enclosing message. The zero value enum should have the suffix `UNSPECIFIED`.
- Note that names of enum entries must be unique in the whole package. Defining a second completely unrelated enum with
  an entry existing in the first will result in an error message because of how enums in C and C++ are implemented. To
  avoid this, prefix the enum entries with the enum name. Some code generators (ex. for C#) will remove these prefixes
  automatically so that the resulting code looks “clean” again.

## Services

If your `.proto` defines an RPC service, you should use `CamelCase` (with an initial capital) for both the service name
and any RPC method names:

```proto
service FooService {
  rpc GetSomething(GetSomethingRequest) returns (GetSomethingResponse);
  rpc ListSomething(ListSomethingRequest) returns (ListSomethingResponse);
}
```

## Separate request and response messages

We recommend that you create a separate message for each request and response. Name them `{MethodName}Request`
and `{MethodName}Reponse`. This allows you to modify request and response messages for a single service method without
introducing accidental changes to other methods. It is tempting to re-use messages and simply ignore fields that aren't
needed. With time, this will result in a mess, since it isn't obvious what the API expects. Exceptions to this rule are
usually made when returning a single, well-defined entity or when returning an empty message.

```proto
service BookService {
  rpc CreateBook(Book) returns (Book); // don't do this
  rpc CreateBook(Book) returns (google.protobuf.Empty); // instead do this
  rpc ListBooks(ListBooksRequest) returns (ListBooksResponse); // this is OK
  rpc GetBook(GetBookRequest) returns (Book); // this is also OK
  rpc DeleteBook(DeleteBookRequest) returns (google.protobuf.Empty); // this is also OK
}
```

## Use the error system

Do not return a simple confirmation to an endpoint that does not need to return any other data,
use `google.protobuf.Empty` and trigger an error if needed.

```proto
// Don't
service BookService {
  rpc DeleteBook(DeleteBookRequest) returns (DeleteBookResponse);
}

message DeleteBookResponse {
  bool deleted = 1;
}

// Do
service BookService {
  rpc DeleteBook(DeleteBookRequest) returns (google.protobuf.Empty);
}
```

## Don’t reinvent the wheel

Use and abuse of the well-known types:

[Package google.protobuf | Protocol Buffers | Google Developers](https://developers.google.com/protocol-buffers/docs/reference/google.protobuf)

Examples:

- The empty message is already defined as google.protobuf.Empty, so it doesn't make sense to define yet another empty
  message.
- Don’t try to redefine timestamp, there’s
  an [existing type for it](https://developers.google.com/protocol-buffers/docs/reference/google.protobuf#timestamp)
- …

## Don’t have one field in a struct influence the meaning of another

Protocol buffer messages usually have multiple fields. These fields should always be independent of each other — you
shouldn’t have one field influence the semantic meaning of another.

```proto
// don't do this!
message Foo {
  int64 timestamp = 1;
  bool timestampIsCreated; // true means timestamp is created time,
                           // false means that it is updated time
```

This causes confusion — the client now needs to have special logic to know how to interpret one field based on another.
Instead, use multiple fields, or use the protobuf “oneof” feature.

```proto
// better, but still not ideal because the fields are mutually
// exclusive - only one will be set
message Foo {
  int64 createdTimestamp;
  int64 updatedTimestamp;
}

// this is ideal; one will be set, and that is enforced by protobuf
message Foo {
  oneof timestamp {
    int64 created;
    int64 updated;
  }
}
```

## Linting

We use protolint to lint the proto files. Please download and
install [protolint](https://github.com/yoheimuta/protolint/releases).

### Git pre-commit hook

You can also enforce protolint to execute automatically before committing code
using [pre-commit](https://pre-commit.com/). Install pre-commit, add a file `.pre-commit-config.yaml`, and add this
content to it:

```yaml
repos:
  - repo: https://github.com/yoheimuta/protolint
    rev: master
    hooks:
      - id: protolint
```

Then run: `pre-commit install` in the repo.

## Extra recommendations

- Use CI to automatically double-check if your proto files still compile without error/warnings.
- Avoid large messages. gRPC is not designed for
  that. [Details](https://kreya.app/blog/grpc-best-practices/#large-messages)
- Reuse channels. Creating a gRPC channel is a costly process, as it creates a new HTTP/2
  connection. [Details](https://kreya.app/blog/grpc-best-practices/#reuse-channels)

## References

- [rotocol-buffers/docs/style](https://developers.google.com/protocol-buffers/docs/style)
- [blog/grpc-best-practices](https://kreya.app/blog/grpc-best-practices/)
- [protobuf-definition-best-practices](https://medium.com/@akhaku/protobuf-definition-best-practices-87f281576f31)
