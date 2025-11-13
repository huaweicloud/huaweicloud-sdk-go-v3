package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteBatchCreateRecordSetsTaskResponse Response Object
type DeleteBatchCreateRecordSetsTaskResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteBatchCreateRecordSetsTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteBatchCreateRecordSetsTaskResponse struct{}"
	}

	return strings.Join([]string{"DeleteBatchCreateRecordSetsTaskResponse", string(data)}, " ")
}
