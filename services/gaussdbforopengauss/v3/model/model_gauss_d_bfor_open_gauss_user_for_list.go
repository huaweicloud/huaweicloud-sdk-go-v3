package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 数据库用户信息。
type GaussDBforOpenGaussUserForList struct {

	// 帐号名。
	Name string `json:"name" xml:"name"`

	Attributes *GaussDBforOpenGaussUserForListAttributes `json:"attributes,omitempty" xml:"attributes"`

	// 用户的默认权限。
	Memberof *string `json:"memberof,omitempty" xml:"memberof"`
}

func (o GaussDBforOpenGaussUserForList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GaussDBforOpenGaussUserForList struct{}"
	}

	return strings.Join([]string{"GaussDBforOpenGaussUserForList", string(data)}, " ")
}
