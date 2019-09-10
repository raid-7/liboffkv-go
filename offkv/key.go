package offkv

import (
	"strings"
	"unicode/utf8"
)

type validKey struct {
	prefix            *string
	data              string
	lastSlashPosition uint64
}

type keyTransformer func(prefix *string, parent string, lastToken string) string

func validateKey(key string) (*validKey, bool) {
	if !strings.HasPrefix(key, "/") || strings.HasSuffix(key, "/") {
		return nil, false
	}

	tokens := strings.Split(key, "/")
	for _, token := range tokens {
		if !isTokenValid(token) {
			return nil, false
		}
	}

	slashPos := len(key) - len(tokens[len(tokens)-1])
	return &validKey{
		nil, key, uint64(slashPos),
	}, true
}

func (key *validKey) withPrefix(prefix *validKey) *validKey {
	key.prefix = &prefix.data
	return key
}

func (key *validKey) transform(transformer keyTransformer) string {
	if key.prefix == nil {
		return ""
	}
	return transformer(key.prefix, key.data[:key.lastSlashPosition], key.data[key.lastSlashPosition+1:])
}

func isTokenValid(token string) bool {
	for t := token; len(t) > 0; {
		r, sz := utf8.DecodeRuneInString(t)

		if sz != 1 || r <= 0x1F || r >= 0x7F {
			return false
		}

		t = t[sz:]
	}

	return len(token) > 0 && token != "." && token != ".." && token != "zookeeper"
}
