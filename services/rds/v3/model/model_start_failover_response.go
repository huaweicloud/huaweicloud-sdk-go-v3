package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type StartFailoverResponse struct {

	// 实例Id
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId"`

	// 节点Id
	NodeId *string `json:"nodeId,omitempty" xml:"nodeId"`

	// 任务Id
	WorkflowId     *string `json:"workflowId,omitempty" xml:"workflowId"`
	HttpStatusCode int     `json:"-"`
}

func (o StartFailoverResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartFailoverResponse struct{}"
	}

	return strings.Join([]string{"StartFailoverResponse", string(data)}, " ")
}
