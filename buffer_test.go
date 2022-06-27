package bufreaderat

import (
	"io"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadAt(t *testing.T) {
	buf := []byte("helloworld")
	b := BufferReaderAt{Buf: buf}

	b1 := make([]byte, len(buf))
	n, err := b.ReadAt(b1, 0)
	require.NoError(t, err)
	require.Equal(t, len(buf), n)
	require.Equal(t, buf, b1)

	n, err = b.ReadAt(b1, 1)
	require.Error(t, err)
	require.ErrorIs(t, err, io.EOF)
	require.Equal(t, len(buf)-1, n)
	require.Equal(t, buf[1:], b1[0:n])

	n, err = b.ReadAt(b1, int64(len(buf)))
	require.Error(t, err)
	require.ErrorIs(t, err, io.EOF)
	require.Equal(t, 0, n)

	n, err = b.ReadAt(b1, int64(len(buf)-1))
	require.Error(t, err)
	require.ErrorIs(t, err, io.EOF)
	require.Equal(t, 1, n)
	require.Equal(t, buf[len(buf)-1:], b1[0:n])

	n, err = b.ReadAt(b1, int64(len(buf)+1))
	require.Error(t, err)
	require.ErrorIs(t, err, ErrInvalidSeek)

	n, err = b.ReadAt(b1, -1)
	require.Error(t, err)
	require.ErrorIs(t, err, ErrInvalidSeek)

	b2 := make([]byte, 0, 4)
	n, err = b.ReadAt(b2, 0)
	require.NoError(t, err)
	require.Equal(t, 0, n)

	b3 := make([]byte, len(buf)*2)
	n, err = b.ReadAt(b3, 0)
	require.Error(t, err)
	require.ErrorIs(t, err, io.EOF)
	require.Equal(t, len(buf), n)
	require.Equal(t, buf, b3[0:len(buf)])
}

func TestSeekRead(t *testing.T) {
	buf := []byte("helloworld")
	b := BufferReaderAt{Buf: buf}

	n, err := b.Seek(0, io.SeekStart)
	require.NoError(t, err)
	require.Equal(t, int64(0), n)

	n, err = b.Seek(int64(len(buf)), io.SeekStart)
	require.NoError(t, err)
	require.Equal(t, int64(len(buf)), n)

	n, err = b.Seek(int64(len(buf)+1), io.SeekStart)
	require.Error(t, err)
	require.ErrorIs(t, err, ErrInvalidSeek)
	require.Equal(t, int64(-1), n)

	n, err = b.Seek(-1, io.SeekStart)
	require.Error(t, err)
	require.ErrorIs(t, err, ErrInvalidSeek)
	require.Equal(t, int64(-1), n)

	b1 := make([]byte, 2*len(buf))
	rd, err := b.Read(b1)
	require.Error(t, err)
	require.ErrorIs(t, err, io.EOF)
	require.Equal(t, 0, rd)

	n, err = b.Seek(int64(0), io.SeekEnd)
	require.NoError(t, err)
	require.Equal(t, int64(len(buf)), n)

	n, err = b.Seek(int64(-len(buf)), io.SeekCurrent)
	require.NoError(t, err)
	require.Equal(t, int64(0), n)

	rd, err = b.Read(b1[0:])
	require.NoError(t, err)
	require.Equal(t, len(buf), rd)
	require.Equal(t, buf, b1[0:rd])

	// small read
	b2 := make([]byte, 2)

	n, err = b.Seek(0, io.SeekStart)
	require.NoError(t, err)
	require.Equal(t, int64(0), n)

	rd, err = b.Read(b2)
	require.NoError(t, err)
	require.Equal(t, buf[0:2], b2)

	n, err = b.Seek(0, io.SeekCurrent)
	require.NoError(t, err)
	require.Equal(t, int64(2), n)

	rd, err = b.Read(b2)
	require.NoError(t, err)
	require.Equal(t, buf[2:4], b2)
}
