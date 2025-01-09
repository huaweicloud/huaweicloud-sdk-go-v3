package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateServerGroupTagsReq 服务器组关联的标签列表。
type BatchCreateServerGroupTagsReq struct {

	// 服务器组关联的标签列表。
	Items []ServerGroupTagsInfo `json:"items"`
}

func (o BatchCreateServerGroupTagsReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateServerGroupTagsReq struct{}"
	}

	return strings.Join([]string{"BatchCreateServerGroupTagsReq", string(data)}, " ")
}
