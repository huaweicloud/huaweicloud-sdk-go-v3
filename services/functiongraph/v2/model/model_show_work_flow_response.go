package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowWorkFlowResponse Response Object
type ShowWorkFlowResponse struct {

	// 唯一标识ID，流程定义ID
	Id *string `json:"id,omitempty"`

	// 函数工作流URN, 格式为： urn:fss:<region_id>:<project_id>:workflow:\\<package\\>:<workflow_name>:\\<version\\> 注意： package当前只支持default version当前只支持latest
	WorkflowUrn *string `json:"workflow_urn,omitempty"`

	// 流程创建时间，格式：yyyy-MM-ddTHH:mm:ssZ，UTC时间
	CreatedTime *string `json:"created_time,omitempty"`

	// 流程修改时间，格式：yyyy-MM-ddTHH:mm:ssZ，UTC时间
	UpdatedTime *string `json:"updated_time,omitempty"`

	// 流程创建者
	CreatedBy *string `json:"created_by,omitempty"`

	// 快速函数流日志组ID，仅快速模式函数流且日志级别不为NONE时
	LtsGroupId *string `json:"lts_group_id,omitempty"`

	// 快速函数流日志流ID，仅快速模式函数流且日志级别不为NONE时返回。
	LtsStreamId *string `json:"lts_stream_id,omitempty"`

	Definition     *WorkflowCreateBody `json:"definition,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ShowWorkFlowResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowWorkFlowResponse struct{}"
	}

	return strings.Join([]string{"ShowWorkFlowResponse", string(data)}, " ")
}
