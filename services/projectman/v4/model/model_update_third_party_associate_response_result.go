package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateThirdPartyAssociateResponseResult 修改工作项下外部链接的响应结果。
type UpdateThirdPartyAssociateResponseResult struct {

	// 修改失败的字段列表。
	Fail *[]string `json:"fail,omitempty"`

	// 成功修改的字段集合，每个元素为一个工作项对应的字段名数组。
	Success *[][]string `json:"success,omitempty"`
}

func (o UpdateThirdPartyAssociateResponseResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateThirdPartyAssociateResponseResult struct{}"
	}

	return strings.Join([]string{"UpdateThirdPartyAssociateResponseResult", string(data)}, " ")
}
