package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ClusterAuthTros struct {

	// 操作界面授权开启或关闭
	Enable bool `json:"enable"`

	// 0为集群界面授权功能的普通授权
	AuthType float32 `json:"auth_type"`

	// enable为true时配置，即授权截止时间，如不传该值默认为7天授权时间
	ExpireTime float32 `json:"expire_time,omitempty"`
}

func (o ClusterAuthTros) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterAuthTros struct{}"
	}

	return strings.Join([]string{"ClusterAuthTros", string(data)}, " ")
}
