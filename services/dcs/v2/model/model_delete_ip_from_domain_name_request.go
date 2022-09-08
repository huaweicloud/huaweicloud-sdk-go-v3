package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteIpFromDomainNameRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// 分片ID。
	GroupId string `json:"group_id"`

	// 节点ID。
	NodeId string `json:"node_id"`
}

func (o DeleteIpFromDomainNameRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteIpFromDomainNameRequest struct{}"
	}

	return strings.Join([]string{"DeleteIpFromDomainNameRequest", string(data)}, " ")
}
