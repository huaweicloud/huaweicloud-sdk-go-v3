package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowUserAccessStagesRequest Request Object
type ShowUserAccessStagesRequest struct {

	// 事务id
	TransactionId string `json:"transaction_id"`
}

func (o ShowUserAccessStagesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowUserAccessStagesRequest struct{}"
	}

	return strings.Join([]string{"ShowUserAccessStagesRequest", string(data)}, " ")
}
