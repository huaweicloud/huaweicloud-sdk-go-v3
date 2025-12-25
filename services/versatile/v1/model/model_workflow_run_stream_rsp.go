package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WorkflowRunStreamRsp 工作流执行接口流式响应体
type WorkflowRunStreamRsp struct {

	// 事件类型：1、workflow_started 2、node_started 3、node_wait 4、node_finished 5、workflow_finished 6、text_chunk ：流式输出到对话框的信息
	Event *string `json:"event,omitempty"`

	// 执行的conversation_id
	ConversationId *string `json:"conversation_id,omitempty"`

	// 类型说明：1. node_started，node_finished，node_wait data使用NodeInfo对象 2. workflow_started, workflow_finished  data使用 WorkflowInfo对象 3. text_chunk使用对象TextChunk
	Data *interface{} `json:"data,omitempty"`

	// 事件发生时间
	CreatedTime *int64 `json:"createdTime,omitempty"`
}

func (o WorkflowRunStreamRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkflowRunStreamRsp struct{}"
	}

	return strings.Join([]string{"WorkflowRunStreamRsp", string(data)}, " ")
}
