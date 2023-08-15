package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AcceptOrRejectEndpointResponse Response Object
type AcceptOrRejectEndpointResponse struct {

	// 连接列表
	Connections    *[]ConnectionEndpoints `json:"connections,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o AcceptOrRejectEndpointResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AcceptOrRejectEndpointResponse struct{}"
	}

	return strings.Join([]string{"AcceptOrRejectEndpointResponse", string(data)}, " ")
}
