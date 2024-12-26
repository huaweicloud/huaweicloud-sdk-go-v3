package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ChannelUpdateReq struct {

	// 通道描述
	Description *string `json:"description,omitempty"`

	// 企业项目id
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 跨账号开关
	CrossAccount *bool `json:"cross_account,omitempty"`

	// 策略
	Policy map[string]ChannelCreateReqPolicy `json:"policy,omitempty"`
}

func (o ChannelUpdateReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChannelUpdateReq struct{}"
	}

	return strings.Join([]string{"ChannelUpdateReq", string(data)}, " ")
}
