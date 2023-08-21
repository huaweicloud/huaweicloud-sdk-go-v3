package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BindHost 绑定的域名信息
type BindHost struct {

	// 域名ID
	Id *string `json:"id,omitempty"`

	// 域名
	Hostname *string `json:"hostname,omitempty"`
}

func (o BindHost) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BindHost struct{}"
	}

	return strings.Join([]string{"BindHost", string(data)}, " ")
}
