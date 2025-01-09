package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteServerGroupTagsReq 服务器组关联的标签列表。
type BatchDeleteServerGroupTagsReq struct {

	// 服务器组关联的标签列表。
	Items []ServerGroupTagsInfo `json:"items"`
}

func (o BatchDeleteServerGroupTagsReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteServerGroupTagsReq struct{}"
	}

	return strings.Join([]string{"BatchDeleteServerGroupTagsReq", string(data)}, " ")
}
