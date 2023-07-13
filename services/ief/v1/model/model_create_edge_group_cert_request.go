package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateEdgeGroupCertRequest Request Object
type CreateEdgeGroupCertRequest struct {

	// 铂金版实例ID，专业版实例为空值
	IefInstanceId *string `json:"ief-instance-id,omitempty"`

	// 边缘节点组ID
	GroupId string `json:"group_id"`

	Body *EdgeGroupCertRequest `json:"body,omitempty"`
}

func (o CreateEdgeGroupCertRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEdgeGroupCertRequest struct{}"
	}

	return strings.Join([]string{"CreateEdgeGroupCertRequest", string(data)}, " ")
}
