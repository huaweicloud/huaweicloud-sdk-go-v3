package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PersistObjectIdDecryptDto struct {

	// 是否加密。 - true：加密。 - false：不加密。
	Decrypt bool `json:"decrypt"`

	// 唯一标识。
	Id string `json:"id"`
}

func (o PersistObjectIdDecryptDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PersistObjectIdDecryptDto struct{}"
	}

	return strings.Join([]string{"PersistObjectIdDecryptDto", string(data)}, " ")
}
