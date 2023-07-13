package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyClusterDns 修改的域名信息。
type ModifyClusterDns struct {

	// 待修改的域名。
	Name string `json:"name"`

	// 域名类型。 - public：公网域名。 - private：内网域名。
	Type string `json:"type"`

	// 用于填写默认生成的SOA记录中有效缓存时间，以秒为单位。 - 取值范围：300~2147483647。 - 默认值为300s。
	Ttl int32 `json:"ttl"`
}

func (o ModifyClusterDns) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyClusterDns struct{}"
	}

	return strings.Join([]string{"ModifyClusterDns", string(data)}, " ")
}
