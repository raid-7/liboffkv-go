package offkv

const (
	ERROR_UNKNOWN             = 0
	ERROR_INVALID_ADDRESS     = 1
	ERROR_INVALID_KEY         = 2
	ERROR_NO_KEY              = 3
	ERROR_KEY_EXISTS          = 4
	ERROR_CHILDREN_FOR_LEASED = 5
	ERROR_CONNECTION_LOSS     = 6
)

var error_messages = map[int32]string{
	ERROR_UNKNOWN:             "Unknown error",
	ERROR_INVALID_ADDRESS:     "Invalid address",
	ERROR_INVALID_KEY:         "Invalid key",
	ERROR_NO_KEY:              "No key",
	ERROR_KEY_EXISTS:          "Key exists",
	ERROR_CHILDREN_FOR_LEASED: "Leased keys cannot have children",
	ERROR_CONNECTION_LOSS:     "Connection lost",
}

type Error interface {
	Error() string
	ErrorCode() int32
}

type offkvError struct {
	code int32
}

func (err offkvError) Error() string {
	return error_messages[err.code]
}

func (err offkvError) ErrorCode() int32 {
	return err.code
}

type offkvCustomError struct {
	error string
	code  int32
}

func (err offkvCustomError) Error() string {
	return err.error
}

func (err offkvCustomError) ErrorCode() int32 {
	return err.code
}
