package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowOperationRequest Request Object
type ShowOperationRequest struct {

	// 操作ID。
	OperationId string `json:"operation_id"`
}

func (o ShowOperationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowOperationRequest struct{}"
	}

	return strings.Join([]string{"ShowOperationRequest", string(data)}, " ")
}
