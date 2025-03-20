package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PolicyGroup struct {

	// 用户组ID，长度为1到64个字符，只包含字母、数字和\"-\"的字符串。
	GroupId string `json:"group_id"`

	// 身份策略的附加时间。
	AttachedAt *sdktime.SdkTime `json:"attached_at"`
}

func (o PolicyGroup) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PolicyGroup struct{}"
	}

	return strings.Join([]string{"PolicyGroup", string(data)}, " ")
}
