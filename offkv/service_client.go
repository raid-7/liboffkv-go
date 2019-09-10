package offkv

type serviceClient interface {
	Create(*validKey, []byte, bool) (uint64, Error)
	Set(*validKey, []byte) (uint64, Error)
	Cas(*validKey, uint64, []byte) (bool, uint64, Error) // return version?
	Delete(*validKey, uint64) (bool, Error)

	Exists(*validKey, bool) (ExistsResult, Error)
	GetValue(*validKey, bool) (GetResult, Error)
	GetChildren(*validKey, bool) (ChildrenResult, Error)

	Commit(Transaction) (TransactionResult, Error)

	Close()
}
