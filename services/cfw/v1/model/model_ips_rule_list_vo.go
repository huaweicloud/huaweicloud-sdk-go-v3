package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IpsRuleListVo struct {
	FwInstanceId *string `json:"fw_instance_id,omitempty"`

	Limit *int32 `json:"limit,omitempty"`

	ObjectId *string `json:"object_id,omitempty"`

	Offset *int32 `json:"offset,omitempty"`

	Records *[]IpsRuleVo `json:"records,omitempty"`

	Total *int32 `json:"total,omitempty"`
}

func (o IpsRuleListVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IpsRuleListVo struct{}"
	}

	return strings.Join([]string{"IpsRuleListVo", string(data)}, " ")
}
