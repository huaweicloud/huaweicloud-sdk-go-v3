package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEcnWithVpcRequest Request Object
type ListEcnWithVpcRequest struct {

	// 企业连接网络ID
	EcnId string `json:"ecn_id"`
}

func (o ListEcnWithVpcRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEcnWithVpcRequest struct{}"
	}

	return strings.Join([]string{"ListEcnWithVpcRequest", string(data)}, " ")
}
