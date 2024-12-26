package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IpsRuleUpdateTimeVo struct {
	IpsType *int32 `json:"ips_type,omitempty"`

	IpsVersion *string `json:"ips_version,omitempty"`

	UpdateTime *int64 `json:"update_time,omitempty"`
}

func (o IpsRuleUpdateTimeVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IpsRuleUpdateTimeVo struct{}"
	}

	return strings.Join([]string{"IpsRuleUpdateTimeVo", string(data)}, " ")
}
