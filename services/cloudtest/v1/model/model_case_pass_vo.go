package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CasePassVo 需求关联用例通过情况
type CasePassVo struct {

	// 需求关联用例通过率
	PassRate *string `json:"pass_rate,omitempty"`

	// 需求关联用例结果与对应的用例数目列表
	ResultNumberList *[]NameAndValueVo `json:"result_number_list,omitempty"`
}

func (o CasePassVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CasePassVo struct{}"
	}

	return strings.Join([]string{"CasePassVo", string(data)}, " ")
}
