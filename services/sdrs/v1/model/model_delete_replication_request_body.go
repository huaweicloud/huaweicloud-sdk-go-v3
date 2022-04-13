package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 删除复制对请求体
type DeleteReplicationRequestBody struct {
	Replication *DeleteReplicationRequestParams `json:"replication"`
}

func (o DeleteReplicationRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteReplicationRequestBody struct{}"
	}

	return strings.Join([]string{"DeleteReplicationRequestBody", string(data)}, " ")
}
