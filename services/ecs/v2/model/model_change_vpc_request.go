package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeVpcRequest Request Object
type ChangeVpcRequest struct {

	// 云服务器ID。
	ServerId string `json:"server_id"`

	Body *ChangeVpcRequestBody `json:"body,omitempty"`
}

func (o ChangeVpcRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeVpcRequest struct{}"
	}

	return strings.Join([]string{"ChangeVpcRequest", string(data)}, " ")
}
