package offkv

import "reflect"

type ExistsResult struct {
	Exists  bool
	Version uint64
	Watch   <-chan interface{}
}

type GetResult struct {
	Value   []byte
	Version uint64
	Watch   <-chan interface{}
}

type ChildrenResult struct {
	Children []string
	Watch    <-chan interface{}
}

func toSingleNilChan(ch interface{}) <-chan interface{} {
	res := make(chan interface{})
	chVal := reflect.ValueOf(ch)

	go func() {
		_, ok := chVal.Recv()
		if ok {
			res <- nil
		}
		close(res)
	}()

	return nil
}
