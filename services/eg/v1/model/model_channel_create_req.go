package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ChannelCreateReq struct {

	// 通道名称，租户下唯一，由字母，数字，点，下划线和中划线组成，必须字母或数字开头，不能是default
	Name string `json:"name"`

	// 通道描述
	Description *string `json:"description,omitempty"`

	// 企业项目id
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 跨账号开关
	CrossAccount *bool `json:"cross_account,omitempty"`

	// 策略
	Policy map[string]ChannelCreateReqPolicy `json:"policy,omitempty"`
}

func (o ChannelCreateReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChannelCreateReq struct{}"
	}

	return strings.Join([]string{"ChannelCreateReq", string(data)}, " ")
}
