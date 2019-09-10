package offkv

import (
	"github.com/control-center/go-zookeeper/zk"
	"time"
)

type ZKService struct {
	conn   *zk.Conn
	events <-chan zk.Event
}

func zooKeyTransformer(prefix *string, parent string, lastToken string) string {
	return *prefix + parent + "/" + lastToken
}

func zooHandleError(err error) Error {
	switch err {
	case nil:
		return nil

	case zk.ErrNodeExists:
		return offkvError{ERROR_KEY_EXISTS}
	case zk.ErrNoNode:
		return offkvError{ERROR_NO_KEY}
	case zk.ErrNoChildrenForEphemerals:
		return offkvError{ERROR_CHILDREN_FOR_LEASED}

	default:
		return offkvCustomError{err.Error(), ERROR_UNKNOWN}
	}
}

func (zoo *ZKService) Cas(*validKey, uint64, []byte) (bool, uint64, Error) {
	panic("implement me")
}

func (zoo *ZKService) Commit(Transaction) (TransactionResult, Error) {
	panic("implement me")
}

func (zoo *ZKService) Create(key *validKey, value []byte, leased bool) (uint64, Error) {
	// TODO

	var flags int32 = 0
	if leased {
		flags = zk.FlagEphemeral
	}

	_, err := zoo.conn.Create(key.transform(zooKeyTransformer), value, flags, nil)
	if err != nil {
		return 0, zooHandleError(err)
	}

	return 1, nil
}

func (zoo *ZKService) Delete(key *validKey, version uint64) (bool, Error) {
	err := zoo.conn.Delete(key.transform(zooKeyTransformer), int32(version))
	// TODO What does it return on version mismatch?
	return true, zooHandleError(err)
}

func (zoo *ZKService) Exists(*validKey, bool) (ExistsResult, Error) {
	panic("implement me")
}

func (zoo *ZKService) GetChildren(*validKey, bool) (ChildrenResult, Error) {
	panic("implement me")
}

func (zoo *ZKService) GetValue(key *validKey, watch bool) (GetResult, Error) {
	var (
		result     GetResult
		stats      *zk.Stat
		err        error
		eventsChan <-chan zk.Event
	)

	if watch {
		result.Value, stats, eventsChan, err = zoo.conn.GetW(key.transform(zooKeyTransformer))
	} else {
		result.Value, stats, err = zoo.conn.Get(key.transform(zooKeyTransformer))
	}

	if err != nil {
		return result, zooHandleError(err)
	}

	result.Version = uint64(stats.Version)
	if watch {
		result.Watch = toSingleNilChan(eventsChan)
	}

	return result, nil
}

func (zoo *ZKService) Set(*validKey, []byte) (uint64, Error) {
	panic("implement me")
}

func (zoo *ZKService) Close() {
	zoo.conn.Close()
}

func newZkService(addresses []string) (*ZKService, Error) {
	var err error
	service := new(ZKService)
	service.conn, service.events, err = zk.Connect(addresses, time.Second*10)

	// TODO init prefix key

	if err != nil {
		return nil, offkvCustomError{err.Error(), ERROR_CONNECTION_LOSS}
	}

	return service, nil
}
