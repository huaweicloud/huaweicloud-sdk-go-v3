package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCentralNetworkRequestBody 创建中心网络的请求体。
type CreateCentralNetworkRequestBody struct {
	CentralNetwork *CreateCentralNetwork `json:"central_network"`
}

func (o CreateCentralNetworkRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCentralNetworkRequestBody struct{}"
	}

	return strings.Join([]string{"CreateCentralNetworkRequestBody", string(data)}, " ")
}
