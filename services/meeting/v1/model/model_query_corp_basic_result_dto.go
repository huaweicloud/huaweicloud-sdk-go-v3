package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 企业基本信息
type QueryCorpBasicResultDto struct {
	// 企业名称

	Name *string `json:"name,omitempty"`
	// 联系号码

	Phone *string `json:"phone,omitempty"`
	// 联系号码所属的国家

	Country *string `json:"country,omitempty"`
	// 传真号码

	Fax *string `json:"fax,omitempty"`
	// 邮箱地址

	Email *string `json:"email,omitempty"`
	// 地址

	Address *string `json:"address,omitempty"`
	// 备注

	Description *string `json:"description,omitempty"`
	// 企业所属spId

	SpId *string `json:"spId,omitempty"`
}

func (o QueryCorpBasicResultDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryCorpBasicResultDto struct{}"
	}

	return strings.Join([]string{"QueryCorpBasicResultDto", string(data)}, " ")
}
