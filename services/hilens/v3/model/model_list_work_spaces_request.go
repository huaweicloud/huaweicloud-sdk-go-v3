package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListWorkSpacesRequest struct {

	// 用户的userId，用于查询指定的的子工作空间
	IamUserId *string `json:"iam_user_id,omitempty"`
}

func (o ListWorkSpacesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWorkSpacesRequest struct{}"
	}

	return strings.Join([]string{"ListWorkSpacesRequest", string(data)}, " ")
}
