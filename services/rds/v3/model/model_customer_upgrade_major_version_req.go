package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CustomerUpgradeMajorVersionReq struct {

	// 是否在运维时间窗内切换版本,默认为否。
	Delay *bool `json:"delay,omitempty"`

	// 指定新版本的参数模板，若不填则默认会最大限度继承原实例的参数。
	ConfigurationId *string `json:"configuration_id,omitempty"`
}

func (o CustomerUpgradeMajorVersionReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CustomerUpgradeMajorVersionReq struct{}"
	}

	return strings.Join([]string{"CustomerUpgradeMajorVersionReq", string(data)}, " ")
}
