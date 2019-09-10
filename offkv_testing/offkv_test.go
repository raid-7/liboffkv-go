package offkv_testing

import (
	"offscale/liboffkv-go/offkv"
	"strings"
	"testing"
)

var ADDRESSES = [...]string{
	"zk://127.0.0.1:2831",
}

var tests = map[string]func(*testing.T, offkv.Client){
	"create_get": func(t *testing.T, cl offkv.Client) {
		defer cl.Delete("mykey")

		createdVersion, err := cl.Create("mykey", []byte("value"))
		failIfError(t, err)

		result, err := cl.Get("mykey", false)
		failIfError(t, err)

		if result.Version != createdVersion {
			t.Error("Version mismatch")
		}

		if string(result.Value) != "value" {
			t.Error("Value mismatch")
		}
	},
	"create_delete": func(t *testing.T, cl offkv.Client) {
		createdVersion, err := cl.Create("mykey", []byte("value"))
		failIfError(t, err)

		_, err = cl.CompareDelete("mykey", createdVersion+1)
		t.Log(err)
	},
}

func TestAPI(t *testing.T) {
	for _, addr := range ADDRESSES {
		serviceName := addr[:strings.Index(addr, "://")]
		t.Run(serviceName, func(t *testing.T) {
			//t.Parallel()  // tests for different services can be run in parallel

			service, err := offkv.New(addr, "/ut")
			if err != nil {
				t.Error(err)
			}
			defer service.Close()

			for name, runner := range tests {
				t.Run(name, func(t *testing.T) {
					runner(t, service)
				})
			}
		})
	}
}

func failIfError(t *testing.T, err offkv.Error) {
	if err != nil {
		t.Error(err)
	}
}
