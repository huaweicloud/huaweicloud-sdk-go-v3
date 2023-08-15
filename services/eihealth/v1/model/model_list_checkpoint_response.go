package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCheckpointResponse Response Object
type ListCheckpointResponse struct {

	// 总条数
	Count *int32 `json:"count,omitempty"`

	// 数据作业执行日志
	Logs           *[]CheckpointRsp `json:"logs,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ListCheckpointResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCheckpointResponse struct{}"
	}

	return strings.Join([]string{"ListCheckpointResponse", string(data)}, " ")
}
