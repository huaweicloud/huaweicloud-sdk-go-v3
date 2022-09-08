package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateNewBlockchainResponse struct {

	// 服务实例ID
	BlockchainId *string `json:"blockchain_id,omitempty"`

	// 服务实例名
	BlockchainName *string `json:"blockchain_name,omitempty"`

	// 操作ID
	OperationId    *string `json:"operation_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateNewBlockchainResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateNewBlockchainResponse struct{}"
	}

	return strings.Join([]string{"CreateNewBlockchainResponse", string(data)}, " ")
}
