package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CasePassRateVo 计算用例通过率
type CasePassRateVo struct {

	// 用例通过率
	PassRate *string `json:"pass_rate,omitempty"`

	// 用户自定义结果对应的用例数目
	ResultNumberList *[]NameAndValueVo `json:"result_number_list,omitempty"`
}

func (o CasePassRateVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CasePassRateVo struct{}"
	}

	return strings.Join([]string{"CasePassRateVo", string(data)}, " ")
}
