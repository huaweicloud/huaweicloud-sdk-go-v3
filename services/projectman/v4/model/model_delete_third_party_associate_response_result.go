package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteThirdPartyAssociateResponseResult 删除工作项下外部链接的响应结果。
type DeleteThirdPartyAssociateResponseResult struct {

	// 删除失败的外部链接ID列表。
	Fail *[]string `json:"fail,omitempty"`

	// 成功删除的外部链接ID列表。
	Success *[]string `json:"success,omitempty"`
}

func (o DeleteThirdPartyAssociateResponseResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteThirdPartyAssociateResponseResult struct{}"
	}

	return strings.Join([]string{"DeleteThirdPartyAssociateResponseResult", string(data)}, " ")
}
