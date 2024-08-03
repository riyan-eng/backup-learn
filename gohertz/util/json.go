package util

import "github.com/bytedance/sonic"

func UnmarshalConverter[T any](s []byte) (data T) {
	sonic.Unmarshal(s, &data)
	return
}
