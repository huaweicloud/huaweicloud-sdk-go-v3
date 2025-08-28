package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportRisksResponse Response Object
type ExportRisksResponse struct {

	// **参数解释**： 任务ID **取值范围**： 字符长度1-64位
	TaskId         *string `json:"task_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ExportRisksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportRisksResponse struct{}"
	}

	return strings.Join([]string{"ExportRisksResponse", string(data)}, " ")
}
