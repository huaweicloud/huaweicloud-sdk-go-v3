package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IdAndStatus **参数说明**：对象的简易状态，包括ID和状态。
type IdAndStatus struct {

	// **参数说明**：摄像头、雷达或RSU的ID。
	Id *string `json:"id,omitempty"`

	// **参数说明**：摄像头、雷达或RSU的状态。  **取值范围**： - ONLINE：在线 - OFFLINE：离线 - INITIAL：初始化
	Status *string `json:"status,omitempty"`
}

func (o IdAndStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IdAndStatus struct{}"
	}

	return strings.Join([]string{"IdAndStatus", string(data)}, " ")
}
