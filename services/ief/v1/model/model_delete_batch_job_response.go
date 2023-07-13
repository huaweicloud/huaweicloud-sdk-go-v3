package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteBatchJobResponse Response Object
type DeleteBatchJobResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteBatchJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteBatchJobResponse struct{}"
	}

	return strings.Join([]string{"DeleteBatchJobResponse", string(data)}, " ")
}
