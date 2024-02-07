package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NativeUuid 资源ID标识符。
type NativeUuid struct {
}

func (o NativeUuid) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NativeUuid struct{}"
	}

	return strings.Join([]string{"NativeUuid", string(data)}, " ")
}
