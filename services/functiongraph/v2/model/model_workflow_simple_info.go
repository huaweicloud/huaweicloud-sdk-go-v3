package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WorkflowSimpleInfo 函数流详情
type WorkflowSimpleInfo struct {

	// 函数流是否返回流式数据
	EnableStreamResponse bool `json:"enable_stream_response"`

	// 唯一标识ID，流程定义ID
	Id string `json:"id"`

	// 函数工作流URN, 格式为： urn:fss:<region_id>:<project_id>:workflow:\\<package\\>:<workflow_name>:\\<version\\> 注意： package当前只支持default version当前只支持latest
	WorkflowUrn string `json:"workflow_urn"`

	// 流程定义名称
	Name string `json:"name"`

	// 流程定义描述
	Description string `json:"description"`

	// 流程创建时间，格式：yyyy-MM-ddTHH:mm:ssZ，UTC时间
	CreatedTime string `json:"created_time"`

	// 流程修改时间，格式：yyyy-MM-ddTHH:mm:ssZ，UTC时间
	UpdatedTime string `json:"updated_time"`

	// 流程创建者
	CreatedBy string `json:"created_by"`
}

func (o WorkflowSimpleInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkflowSimpleInfo struct{}"
	}

	return strings.Join([]string{"WorkflowSimpleInfo", string(data)}, " ")
}
