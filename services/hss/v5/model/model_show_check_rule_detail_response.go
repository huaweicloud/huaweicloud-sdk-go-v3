package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowCheckRuleDetailResponse struct {

	// 描述
	Description *string `json:"description,omitempty" xml:"description"`

	// 根据
	Reference *string `json:"reference,omitempty" xml:"reference"`

	// 审计描述
	Audit *string `json:"audit,omitempty" xml:"audit"`

	// 修改建议
	Remediation *string `json:"remediation,omitempty" xml:"remediation"`

	// 检测用例信息
	CheckInfoList  *[]CheckRuleCheckCaseResponseInfo `json:"check_info_list,omitempty" xml:"check_info_list"`
	HttpStatusCode int                               `json:"-"`
}

func (o ShowCheckRuleDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCheckRuleDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowCheckRuleDetailResponse", string(data)}, " ")
}
