package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TagsReq struct {

	// 标签列表。租户权限时该字段必选，op_service权限时和sys_tags二选一。
	Tags *[]ResourceTag `json:"tags,omitempty"`
}

func (o TagsReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagsReq struct{}"
	}

	return strings.Join([]string{"TagsReq", string(data)}, " ")
}
