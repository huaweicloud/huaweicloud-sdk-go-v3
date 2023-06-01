package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MetadataDeleteReq struct {

	// 需要删除的任务列表。
	TaskIds *[]string `json:"taskIds,omitempty"`
}

func (o MetadataDeleteReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MetadataDeleteReq struct{}"
	}

	return strings.Join([]string{"MetadataDeleteReq", string(data)}, " ")
}
