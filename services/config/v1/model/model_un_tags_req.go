package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UnTagsReq struct {

	// 标签列表。租户权限时该字段必选，op_service权限时和sys_tags二选一。
	Tags *[]ResourceUnTag `json:"tags,omitempty"`
}

func (o UnTagsReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UnTagsReq struct{}"
	}

	return strings.Join([]string{"UnTagsReq", string(data)}, " ")
}
