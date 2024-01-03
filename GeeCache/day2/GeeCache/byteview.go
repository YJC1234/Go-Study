package geecache

// ByteView holds a immutable view of bytes
type ByteView struct {
	b []byte
}

// Len 实现Len接口
func (v ByteView) Len() int {
	return len(v.b)
}

// ByteSlice copy
func (v ByteView) ByteSlice() []byte {
	return cloneBytes(v.b)
}

func (v ByteView) String() string {
	return string(v.b)
}

func cloneBytes(b []byte) []byte {
	c := make([]byte, len(b))
	copy(c, b)
	return c
}
