package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowIpdThirdPartyAssociatedRequest Request Object
type ShowIpdThirdPartyAssociatedRequest struct {

	// 工作项唯一ID。可以通过查询工作项列表或者查询树状工作项获取，响应消息体中的ID字段的值就是工作项ID。
	IssueId string `json:"issue_id"`

	// 项目32位ID，项目唯一标识，通过查询IPD项目列表获取，响应消息体中的project_id字段的值就是项目ID。
	ProjectId string `json:"project_id"`
}

func (o ShowIpdThirdPartyAssociatedRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowIpdThirdPartyAssociatedRequest struct{}"
	}

	return strings.Join([]string{"ShowIpdThirdPartyAssociatedRequest", string(data)}, " ")
}
