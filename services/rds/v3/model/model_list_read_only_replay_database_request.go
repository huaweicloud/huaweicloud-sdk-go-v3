package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListReadOnlyReplayDatabaseRequest Request Object
type ListReadOnlyReplayDatabaseRequest struct {

	// 实例id
	InstanceId string `json:"instance_id"`
}

func (o ListReadOnlyReplayDatabaseRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListReadOnlyReplayDatabaseRequest struct{}"
	}

	return strings.Join([]string{"ListReadOnlyReplayDatabaseRequest", string(data)}, " ")
}
