package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateVectorStorageConfigReq struct {

	// **参数解释**：  向量存储开关。true表示开启向量存储，false表示关闭向量存储。  **约束限制**：  不涉及。  **取值范围**：  - true - false **默认取值**：  不涉及。
	VectorStorageSwitch bool `json:"vector_storage_switch"`
}

func (o UpdateVectorStorageConfigReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateVectorStorageConfigReq struct{}"
	}

	return strings.Join([]string{"UpdateVectorStorageConfigReq", string(data)}, " ")
}
