package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowWorkFlowResponse struct {

	// 唯一标识ID，流程定义ID
	Id *string `json:"id,omitempty"`

	// 唯一标识ID，流程URN
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

	Definition     *CreateWorkflowRequestBody `json:"definition,omitempty"`
	HttpStatusCode int                        `json:"-"`
}

func (o ShowWorkFlowResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowWorkFlowResponse struct{}"
	}

	return strings.Join([]string{"ShowWorkFlowResponse", string(data)}, " ")
}
