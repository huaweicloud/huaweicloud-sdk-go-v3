package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateNetworkInstanceResponse Response Object
type CreateNetworkInstanceResponse struct {

	// 请求ID。
	RequestId string `json:"request_id"`

	NetworkInstance *NetworkInstance `json:"network_instance"`
	HttpStatusCode  int              `json:"-"`
}

func (o CreateNetworkInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateNetworkInstanceResponse struct{}"
	}

	return strings.Join([]string{"CreateNetworkInstanceResponse", string(data)}, " ")
}
