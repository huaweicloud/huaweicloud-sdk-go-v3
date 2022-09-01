package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Features struct {

	// 是否支持acl - true：是 - false：否
	SupportAcl *bool `json:"support_acl,omitempty" xml:"support_acl"`

	// 实例是否支持客户端ip透传 - true：是 - false：否
	SupportTransparentClientIp *bool `json:"support_transparent_client_ip,omitempty" xml:"support_transparent_client_ip"`

	// 是否支持SSL - true：是 - false：否
	SupportSsl *bool `json:"support_ssl,omitempty" xml:"support_ssl"`
}

func (o Features) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Features struct{}"
	}

	return strings.Join([]string{"Features", string(data)}, " ")
}
