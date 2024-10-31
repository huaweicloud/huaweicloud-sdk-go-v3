package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteHttpReferenceTableRequest Request Object
type DeleteHttpReferenceTableRequest struct {

	// 引用表id
	TableId string `json:"table_id"`
}

func (o DeleteHttpReferenceTableRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteHttpReferenceTableRequest struct{}"
	}

	return strings.Join([]string{"DeleteHttpReferenceTableRequest", string(data)}, " ")
}
