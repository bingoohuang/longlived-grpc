# Google Protocol Buffer

This module contains articles about Google Protocol Buffer.

`protoc --java_out=src/main/java/ src/main/resources/addressbook.proto`

## Relevant articles

- [Introduction to Google Protocol Buffer](https://www.baeldung.com/google-protocol-buffer)

## Parsing and Serialization

Finally, each protocol buffer class has methods for writing and reading messages of your chosen type using the protocol buffer binary format. These include:

- `byte[] toByteArray();`: serializes the message and returns a byte array containing its raw bytes.
- `static Person parseFrom(byte[] data);`: parses a message from the given byte array.
- `void writeTo(OutputStream output);`: serializes the message and writes it to an OutputStream.
- `static Person parseFrom(InputStream input);`: reads and parses a message from an InputStream.

