package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPersonnelRequest Request Object
type ListPersonnelRequest struct {

	// 是否有手机号
	HasMobile *bool `json:"has_mobile,omitempty"`

	// IAM账号
	Name *string `json:"name,omitempty"`

	// 偏移量
	Offset int64 `json:"offset"`

	// 分页
	Limit int32 `json:"limit"`
}

func (o ListPersonnelRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPersonnelRequest struct{}"
	}

	return strings.Join([]string{"ListPersonnelRequest", string(data)}, " ")
}
