package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateEdgeGroupRequest Request Object
type CreateEdgeGroupRequest struct {

	// 铂金版实例ID，专业版实例为空值
	IefInstanceId *string `json:"ief-instance-id,omitempty"`

	Body *EdgeGroupRequest `json:"body,omitempty"`
}

func (o CreateEdgeGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEdgeGroupRequest struct{}"
	}

	return strings.Join([]string{"CreateEdgeGroupRequest", string(data)}, " ")
}
