package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RdsErrorLogDownload struct {

	// **参数解释**：  任务ID。  **取值范围**：  不涉及。
	JobId *string `json:"job_id,omitempty"`

	// **参数解释**：  日志文件名称。  **取值范围**：  不涉及。
	FileName *string `json:"file_name,omitempty"`

	// **参数解释**：  状态。  **取值范围**：  不涉及。
	Status *string `json:"status,omitempty"`

	// **参数解释**：  文件大小。  **取值范围**：  不涉及。
	FileSize *string `json:"file_size,omitempty"`

	// **参数解释**：  文件下载链接。  **取值范围**：  不涉及。
	FileLink *string `json:"file_link,omitempty"`

	// **参数解释**：  创建时间。  **取值范围**：  不涉及。
	CreateAt *int64 `json:"create_at,omitempty"`

	// **参数解释**：  更新时间。  **取值范围**：  不涉及。
	UpdatedAt *int64 `json:"updated_at,omitempty"`
}

func (o RdsErrorLogDownload) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RdsErrorLogDownload struct{}"
	}

	return strings.Join([]string{"RdsErrorLogDownload", string(data)}, " ")
}
