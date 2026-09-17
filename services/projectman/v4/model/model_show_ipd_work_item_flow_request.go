package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowIpdWorkItemFlowRequest Request Object
type ShowIpdWorkItemFlowRequest struct {

	// 项目32位ID，项目唯一标识。通过查询IPD项目列表获取，响应消息体中的id字段的值就是项目ID。
	ProjectId string `json:"project_id"`

	// 工作项唯一ID。可以通过查询工作项列表或者查询树状工作项接口获取，响应消息体中的id字段的值就是工作项ID。
	IssueId string `json:"issue_id"`

	// **参数解释**： 工作项类型。 **约束限制**： 不涉及。 **取值范围**： 支持多种工作项类型，使用英文逗号分隔，例如：category=IR,SR,AR。 - 系统设备类项目：RR、SF、IR、SR、AR、Task、Bug - 独立软件类项目：RR、SF、IR、US、Task、Bug - 云服务类项目：RR、Epic、FE、US、Task、Bug **默认取值**： 不涉及。
	IssueCategory string `json:"issue_category"`
}

func (o ShowIpdWorkItemFlowRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowIpdWorkItemFlowRequest struct{}"
	}

	return strings.Join([]string{"ShowIpdWorkItemFlowRequest", string(data)}, " ")
}
