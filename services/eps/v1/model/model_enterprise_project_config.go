package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EnterpriseProjectConfig struct {

	// 是否支持企业项目删除
	DeleteEpSupport *string `json:"delete_ep_support,omitempty"`
}

func (o EnterpriseProjectConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnterpriseProjectConfig struct{}"
	}

	return strings.Join([]string{"EnterpriseProjectConfig", string(data)}, " ")
}
