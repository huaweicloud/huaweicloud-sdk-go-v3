package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PersistObjectIdsDecryptDto struct {

	// 是否加密。 - true：加密。 - false：不加密。
	Decrypt *bool `json:"decrypt,omitempty"`

	// ID列表。
	Ids []int64 `json:"ids"`
}

func (o PersistObjectIdsDecryptDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PersistObjectIdsDecryptDto struct{}"
	}

	return strings.Join([]string{"PersistObjectIdsDecryptDto", string(data)}, " ")
}
