package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UrlDto struct {

	// url路径名称
	Name *string `json:"name,omitempty"`

	// 父权限集是否包含此权限, true包含, false不包含, null未检测
	Contains *bool `json:"contains,omitempty"`
}

func (o UrlDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UrlDto struct{}"
	}

	return strings.Join([]string{"UrlDto", string(data)}, " ")
}
