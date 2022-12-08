package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RestartGaussMySqlNodeRequest struct {

	// 语言。
	XLanguage *string `json:"X-Language,omitempty"`

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// 节点ID。
	NodeId string `json:"node_id"`

	Body *RestartNodeRequest `json:"body,omitempty"`
}

func (o RestartGaussMySqlNodeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestartGaussMySqlNodeRequest struct{}"
	}

	return strings.Join([]string{"RestartGaussMySqlNodeRequest", string(data)}, " ")
}
