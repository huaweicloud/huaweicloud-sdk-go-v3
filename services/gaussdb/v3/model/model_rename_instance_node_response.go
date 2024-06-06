package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RenameInstanceNodeResponse Response Object
type RenameInstanceNodeResponse struct {

	// 工作流ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RenameInstanceNodeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RenameInstanceNodeResponse struct{}"
	}

	return strings.Join([]string{"RenameInstanceNodeResponse", string(data)}, " ")
}
