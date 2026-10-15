package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowVectorStorageConfigResponse Response Object
type ShowVectorStorageConfigResponse struct {

	// **参数解释**：  向量存储开关。true表示已开启向量存储，false表示已关闭向量存储。  **约束限制**：  不涉及。  **取值范围**：  - true - false **默认取值**：  不涉及。
	VectorStorageSwitch *bool `json:"vector_storage_switch,omitempty"`
	HttpStatusCode      int   `json:"-"`
}

func (o ShowVectorStorageConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVectorStorageConfigResponse struct{}"
	}

	return strings.Join([]string{"ShowVectorStorageConfigResponse", string(data)}, " ")
}
