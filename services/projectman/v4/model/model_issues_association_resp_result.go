package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IssuesAssociationRespResult 给工作项关联外部链接的响应结果。
type IssuesAssociationRespResult struct {

	// 关联失败的工作项列表。
	Fail *[]CreateThirdPartyAssociateDto `json:"fail,omitempty"`

	// 关联成功的工作项列表。
	Success *[]CreateThirdPartyAssociateDto `json:"success,omitempty"`
}

func (o IssuesAssociationRespResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IssuesAssociationRespResult struct{}"
	}

	return strings.Join([]string{"IssuesAssociationRespResult", string(data)}, " ")
}
