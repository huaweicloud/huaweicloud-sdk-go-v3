package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EventForensicInfoAbnormalLoginForensic **参数解释**： 异常登录取证信息 **取值范围**： 不涉及
type EventForensicInfoAbnormalLoginForensic struct {

	// **参数解释**： IP **取值范围**： 不涉及
	Ip *string `json:"ip,omitempty"`

	// **参数解释**： 用户 **取值范围**： 不涉及
	User *string `json:"user,omitempty"`

	// **参数解释**： 国家 **取值范围**： 不涉及
	Country *string `json:"country,omitempty"`

	// **参数解释**： 省份 **取值范围**： 不涉及
	SubDivision *string `json:"sub_division,omitempty"`

	// **参数解释**： 城市 **取值范围**： 不涉及
	City *string `json:"city,omitempty"`

	// **参数解释**： 登录源（映射地名） **取值范围**： 不涉及
	CityId *int32 `json:"city_id,omitempty"`
}

func (o EventForensicInfoAbnormalLoginForensic) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EventForensicInfoAbnormalLoginForensic struct{}"
	}

	return strings.Join([]string{"EventForensicInfoAbnormalLoginForensic", string(data)}, " ")
}
