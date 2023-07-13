package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GaussDBforOpenGaussUserForList 数据库用户信息。
type GaussDBforOpenGaussUserForList struct {

	// 帐号名。
	Name string `json:"name"`

	Attributes *GaussDBforOpenGaussUserForListAttributes `json:"attributes,omitempty"`

	// 用户的默认权限。
	Memberof *string `json:"memberof,omitempty"`
}

func (o GaussDBforOpenGaussUserForList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GaussDBforOpenGaussUserForList struct{}"
	}

	return strings.Join([]string{"GaussDBforOpenGaussUserForList", string(data)}, " ")
}
