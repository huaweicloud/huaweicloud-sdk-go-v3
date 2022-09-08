package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateWorkflowResponse struct {

	// 唯一标识ID，流程定义ID
	Id *string `json:"id,omitempty"`

	// 函数工作流URN, 格式为： urn:fss:<region_id>:<project_id>:workflow:\\<package\\>:<workflow_name>:\\<version\\> 注意： package当前只支持default version当前只支持latest
	WorkflowUrn *string `json:"workflow_urn,omitempty"`

	// 流程定义名称
	Name *string `json:"name,omitempty"`

	// 流程定义描述
	Description *string `json:"description,omitempty"`

	// 流程创建时间，格式：yyyy-MM-ddTHH:mm:ssZ，UTC时间
	CreatedTime *string `json:"created_time,omitempty"`

	// 流程修改时间，格式：yyyy-MM-ddTHH:mm:ssZ，UTC时间
	UpdatedTime *string `json:"updated_time,omitempty"`

	// 流程创建者
	CreatedBy      *string `json:"created_by,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateWorkflowResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateWorkflowResponse struct{}"
	}

	return strings.Join([]string{"CreateWorkflowResponse", string(data)}, " ")
}
