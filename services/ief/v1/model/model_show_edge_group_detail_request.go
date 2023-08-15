package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowEdgeGroupDetailRequest Request Object
type ShowEdgeGroupDetailRequest struct {

	// 铂金版实例ID，专业版实例为空值
	IefInstanceId *string `json:"ief-instance-id,omitempty"`

	// 边缘节点组ID
	GroupId string `json:"group_id"`
}

func (o ShowEdgeGroupDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEdgeGroupDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowEdgeGroupDetailRequest", string(data)}, " ")
}
