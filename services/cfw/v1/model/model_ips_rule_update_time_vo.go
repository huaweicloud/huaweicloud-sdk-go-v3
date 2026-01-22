package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IpsRuleUpdateTimeVo struct {

	// ips类型，0表示基础防御，1表示虚拟补丁
	IpsType *int32 `json:"ips_type,omitempty"`

	// ips规则版本
	IpsVersion *string `json:"ips_version,omitempty"`

	// ips更新时间戳
	UpdateTime *int64 `json:"update_time,omitempty"`
}

func (o IpsRuleUpdateTimeVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IpsRuleUpdateTimeVo struct{}"
	}

	return strings.Join([]string{"IpsRuleUpdateTimeVo", string(data)}, " ")
}
