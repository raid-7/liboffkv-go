package offkv

import (
	"strings"
)

type Client struct {
	prefix  *validKey
	service serviceClient
}

func (cl Client) Create(key string, value []byte) (uint64, Error) {
	vKey, ok := validateKey(key)
	if !ok {
		return 0, offkvError{ERROR_INVALID_KEY}
	}

	return cl.service.Create(vKey.withPrefix(cl.prefix), value, false)
}

func (cl Client) CreateLeased(key string, value []byte) (uint64, Error) {
	vKey, ok := validateKey(key)
	if !ok {
		return 0, offkvError{ERROR_INVALID_KEY}
	}

	return cl.service.Create(vKey.withPrefix(cl.prefix), value, true)
}

func (cl Client) Set(key string, value []byte) (uint64, Error) {
	vKey, ok := validateKey(key)
	if !ok {
		return 0, offkvError{ERROR_INVALID_KEY}
	}

	return cl.service.Set(vKey.withPrefix(cl.prefix), value)
}

func (cl Client) Cas(key string, expectedVersion uint64, value []byte) (success bool, version uint64, error Error) {
	vKey, ok := validateKey(key)
	if !ok {
		return false, 0, offkvError{ERROR_INVALID_KEY}
	}

	return cl.service.Cas(vKey.withPrefix(cl.prefix), expectedVersion, value)
}

func (cl Client) Delete(key string) (error Error) {
	vKey, ok := validateKey(key)
	if !ok {
		return offkvError{ERROR_INVALID_KEY}
	}

	_, err := cl.service.Delete(vKey.withPrefix(cl.prefix), 0)
	return err
}

func (cl Client) CompareDelete(key string, expectedVersion uint64) (bool, Error) {
	vKey, ok := validateKey(key)
	if !ok {
		return false, offkvError{ERROR_INVALID_KEY}
	}

	return cl.service.Delete(vKey.withPrefix(cl.prefix), expectedVersion)
}

func (cl Client) Exists(key string, watch bool) (ExistsResult, Error) {
	vKey, ok := validateKey(key)
	if !ok {
		return ExistsResult{}, offkvError{ERROR_INVALID_KEY}
	}

	return cl.service.Exists(vKey.withPrefix(cl.prefix), watch)
}

func (cl Client) Get(key string, watch bool) (GetResult, Error) {
	vKey, ok := validateKey(key)
	if !ok {
		return GetResult{}, offkvError{ERROR_INVALID_KEY}
	}

	return cl.service.GetValue(vKey.withPrefix(cl.prefix), watch)
}

func (cl Client) GetChildren(key string, watch bool) (ChildrenResult, Error) {
	vKey, ok := validateKey(key)
	if !ok {
		return ChildrenResult{}, offkvError{ERROR_INVALID_KEY}
	}

	return cl.service.GetChildren(vKey.withPrefix(cl.prefix), watch)
}

func (cl Client) Commit(transaction Transaction) (TransactionResult, Error) {
	return cl.service.Commit(transaction)
}

func (cl Client) Close() {
	cl.service.Close()
}

func New(address string, prefix string) (Client, Error) {
	var (
		service serviceClient
		err     Error
	)

	vPrefix, ok := validateKey(prefix)
	if !ok {
		return Client{}, offkvError{ERROR_INVALID_KEY}
	}

	switch getProtocol(address) {
	case "zk":
		service, err = newZkService([]string{address})
	case "consul":
		return Client{}, nil
	case "etcd":
		return Client{}, nil

	default:
		return Client{}, offkvError{ERROR_INVALID_ADDRESS}
	}

	return Client{vPrefix, service}, err
}

func getProtocol(addr string) string {
	pos := strings.Index(addr, "://")
	if pos < 0 {
		return ""
	}

	return addr[:pos]
}
