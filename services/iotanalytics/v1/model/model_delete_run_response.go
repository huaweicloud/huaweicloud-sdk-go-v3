package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteRunResponse struct {

	// 被取消作业运行ID
	RunId          *string `json:"run_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteRunResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteRunResponse struct{}"
	}

	return strings.Join([]string{"DeleteRunResponse", string(data)}, " ")
}
