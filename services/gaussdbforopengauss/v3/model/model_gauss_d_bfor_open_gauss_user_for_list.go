package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GaussDBforOpenGaussUserForList 数据库用户信息。
type GaussDBforOpenGaussUserForList struct {

	// 帐号名。
	Name string `json:"name"`

	Attribute *GaussDBforOpenGaussUserForListAttribute `json:"attribute,omitempty"`

	// 用户的默认权限。
	Memberof *string `json:"memberof,omitempty"`

	// 用户是否被锁，取值为“true”或“false”。
	LockStatus *bool `json:"lock_status,omitempty"`
}

func (o GaussDBforOpenGaussUserForList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GaussDBforOpenGaussUserForList struct{}"
	}

	return strings.Join([]string{"GaussDBforOpenGaussUserForList", string(data)}, " ")
}
