package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListWorkSpacesRequest Request Object
type ListWorkSpacesRequest struct {

	// 用户的userId，用于查询指定的的子工作空间
	IamUserId *string `json:"iam_user_id,omitempty"`

	// 查询的起始位置，取值范围为非负整数，默认为0
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示的条目数量，取值范围1~100，默认为100
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListWorkSpacesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWorkSpacesRequest struct{}"
	}

	return strings.Join([]string{"ListWorkSpacesRequest", string(data)}, " ")
}
