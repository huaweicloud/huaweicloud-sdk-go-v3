package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDbUsersRequest Request Object
type ListDbUsersRequest struct {

	// 语言
	XLanguage *string `json:"X-Language,omitempty"`

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// 偏移量表示从此偏移量开始查询, offset大于等于0。
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示的条目数量,取值范围[1, 100]。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListDbUsersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDbUsersRequest struct{}"
	}

	return strings.Join([]string{"ListDbUsersRequest", string(data)}, " ")
}
