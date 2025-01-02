package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestoreLogReplayDatabaseRequest Request Object
type RestoreLogReplayDatabaseRequest struct {

	// 实例id
	InstanceId string `json:"instance_id"`

	// 语言
	XLanguage string `json:"X-Language"`

	Body *LogReplayDatabaseReq `json:"body,omitempty"`
}

func (o RestoreLogReplayDatabaseRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreLogReplayDatabaseRequest struct{}"
	}

	return strings.Join([]string{"RestoreLogReplayDatabaseRequest", string(data)}, " ")
}
