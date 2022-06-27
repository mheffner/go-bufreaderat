# go-bufreaderat

This library provides a zero-copy implementation of the [`io.ReaderAt`](https://pkg.go.dev/io#ReaderAt)
and [`io.ReadSeeker`](https://pkg.go.dev/io#ReadSeeker)
interfaces for a `byte[]` slice. For example, this can provide an efficient method when
uploading in-memory buffers to
[AWS S3](https://aws.github.io/aws-sdk-go-v2/docs/sdk-utilities/s3/#putobjectinput-body-field-ioreadseeker-vs-ioreader).

## Use

Grab the latest version:

```
go get github.com/mheffner/go-bufreaderat
```

Example:

```go
	buf := []byte("helloworld")
	b := BufferReaderAt{Buf: buf}

	var _ io.ReaderAt = &b
	var _ io.ReadSeeker = &b
```

The method `ReadAt` is safe for concurrent use and can be invoked from multiple
goroutines with different output buffers and offsets.

The methods `Read` and `Seek` are not safe to execute concurrently since they rely on an
offset pointer.

**NOTE: The underlying buffer must not be changed while the BufferReaderAt is used.**


