package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateTransactionSplitStatusResponse struct {

	// 工作流ID
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateTransactionSplitStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTransactionSplitStatusResponse struct{}"
	}

	return strings.Join([]string{"UpdateTransactionSplitStatusResponse", string(data)}, " ")
}
