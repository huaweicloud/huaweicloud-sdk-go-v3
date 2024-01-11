package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEcnWithVpcResponse Response Object
type ListEcnWithVpcResponse struct {

	// 企业连接网络与虚拟私有云的绑定关系列表
	EcnVpcRelationships *[]EcnWithVpcItem `json:"ecn_vpc_relationships,omitempty"`
	HttpStatusCode      int               `json:"-"`
}

func (o ListEcnWithVpcResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEcnWithVpcResponse struct{}"
	}

	return strings.Join([]string{"ListEcnWithVpcResponse", string(data)}, " ")
}
