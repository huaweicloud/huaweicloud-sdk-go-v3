package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowEdgeGroupCertDetailRequest Request Object
type ShowEdgeGroupCertDetailRequest struct {

	// 铂金版实例ID，专业版实例为空值
	IefInstanceId *string `json:"ief-instance-id,omitempty"`

	// 边缘节点组ID
	GroupId string `json:"group_id"`

	// 边缘节点组证书ID
	GroupCertId string `json:"group_cert_id"`
}

func (o ShowEdgeGroupCertDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEdgeGroupCertDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowEdgeGroupCertDetailRequest", string(data)}, " ")
}
