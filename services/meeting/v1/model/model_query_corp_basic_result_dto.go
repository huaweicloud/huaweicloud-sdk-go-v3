package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 企业基本信息
type QueryCorpBasicResultDto struct {

	// 企业名称
	Name *string `json:"name,omitempty" xml:"name"`

	// 联系号码
	Phone *string `json:"phone,omitempty" xml:"phone"`

	// 联系号码所属的国家
	Country *string `json:"country,omitempty" xml:"country"`

	// 传真号码
	Fax *string `json:"fax,omitempty" xml:"fax"`

	// 邮箱地址
	Email *string `json:"email,omitempty" xml:"email"`

	// 地址
	Address *string `json:"address,omitempty" xml:"address"`

	// 备注
	Description *string `json:"description,omitempty" xml:"description"`

	// 企业所属spId
	SpId *string `json:"spId,omitempty" xml:"spId"`
}

func (o QueryCorpBasicResultDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryCorpBasicResultDto struct{}"
	}

	return strings.Join([]string{"QueryCorpBasicResultDto", string(data)}, " ")
}
