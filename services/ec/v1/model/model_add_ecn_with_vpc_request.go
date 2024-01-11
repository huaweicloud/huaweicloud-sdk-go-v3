package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddEcnWithVpcRequest Request Object
type AddEcnWithVpcRequest struct {

	// 企业连接网络ID
	EcnId string `json:"ecn_id"`

	Body *AddEcnWithVpcRequestBody `json:"body,omitempty"`
}

func (o AddEcnWithVpcRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddEcnWithVpcRequest struct{}"
	}

	return strings.Join([]string{"AddEcnWithVpcRequest", string(data)}, " ")
}
