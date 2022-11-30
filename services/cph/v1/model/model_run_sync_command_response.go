package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type RunSyncCommandResponse struct {

	// 请求的唯一标识ID。
	RequestId *string `json:"request_id,omitempty"`

	// 任务列表。
	Jobs           *[]RunSyncCommandJob `json:"jobs,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o RunSyncCommandResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunSyncCommandResponse struct{}"
	}

	return strings.Join([]string{"RunSyncCommandResponse", string(data)}, " ")
}
