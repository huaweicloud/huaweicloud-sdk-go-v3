package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAppGroupAuthorizationRequest Request Object
type ListAppGroupAuthorizationRequest struct {

	// 单次查询的大小[1-100]
	Limit *int32 `json:"limit,omitempty"`

	// 查询的偏移量
	Offset *int32 `json:"offset,omitempty"`

	// 应用组ID
	AppGroupId *string `json:"app_group_id,omitempty"`

	// 应用授权的用户(组)，精确查询。
	Account *string `json:"account,omitempty"`
}

func (o ListAppGroupAuthorizationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAppGroupAuthorizationRequest struct{}"
	}

	return strings.Join([]string{"ListAppGroupAuthorizationRequest", string(data)}, " ")
}
