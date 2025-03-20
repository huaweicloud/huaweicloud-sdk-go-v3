package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PolicyAgency struct {

	// 委托或信任委托ID，长度为1到64个字符，只包含字母、数字和\"-\"的字符串。
	AgencyId string `json:"agency_id"`

	// 身份策略的附加时间。
	AttachedAt *sdktime.SdkTime `json:"attached_at"`
}

func (o PolicyAgency) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PolicyAgency struct{}"
	}

	return strings.Join([]string{"PolicyAgency", string(data)}, " ")
}
