package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DdlLogInfo struct {

	// **参数解释**：  日志文件ID。  **取值范围**：  不涉及。
	Id string `json:"id"`

	// **参数解释**：  日志文件名称。  **取值范围**：  不涉及。
	FileName string `json:"file_name"`

	// **参数解释**：  日志文件大小，单位为字节。  **取值范围**：  不涉及。
	FileSize int32 `json:"file_size"`

	// **参数解释**：  日志文件上传的创建时间。  **取值范围**：  不涉及。
	CreateTime string `json:"create_time"`

	// **参数解释**：  日志文件上传的结束时间。  **取值范围**：  不涉及。
	EndTime string `json:"end_time"`

	// **参数解释**：  日志文件的状态。  **取值范围**：  - Active：表示正常。 - Disable：表示不可用。
	Status string `json:"status"`
}

func (o DdlLogInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DdlLogInfo struct{}"
	}

	return strings.Join([]string{"DdlLogInfo", string(data)}, " ")
}
