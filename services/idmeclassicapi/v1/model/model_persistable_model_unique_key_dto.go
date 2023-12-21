package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PersistableModelUniqueKeyDto struct {

	// 是否加密。 - true：加密。 - false：不加密。
	Decrypt *bool `json:"decrypt,omitempty"`

	// 示例模型的唯一键属性。
	UniqueKey string `json:"uniqueKey"`
}

func (o PersistableModelUniqueKeyDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PersistableModelUniqueKeyDto struct{}"
	}

	return strings.Join([]string{"PersistableModelUniqueKeyDto", string(data)}, " ")
}
