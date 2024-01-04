package geecache

// PeerPicker 接口，实现PickPeer，根据key选择对应的节点PeerGetter
type PeerPicker interface {
	PickPeer(key string) (PeerGetter, bool)
}

// PeerGetter 接口，实现Get,从group和key查找对应的缓存值
type PeerGetter interface {
	Get(group string, key string) ([]byte, error)
}
