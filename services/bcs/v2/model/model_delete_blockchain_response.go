package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteBlockchainResponse struct {

	// 操作记录id
	OperationId    *string `json:"operation_id,omitempty" xml:"operation_id"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteBlockchainResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteBlockchainResponse struct{}"
	}

	return strings.Join([]string{"DeleteBlockchainResponse", string(data)}, " ")
}
