package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteEdgeGroupCertRequest struct {

	// 铂金版实例ID，专业版实例为空值
	IefInstanceId *string `json:"ief-instance-id,omitempty"`

	// 边缘节点组ID
	GroupId string `json:"group_id"`

	// 边缘节点组证书ID
	GroupCertId string `json:"group_cert_id"`
}

func (o DeleteEdgeGroupCertRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteEdgeGroupCertRequest struct{}"
	}

	return strings.Join([]string{"DeleteEdgeGroupCertRequest", string(data)}, " ")
}
