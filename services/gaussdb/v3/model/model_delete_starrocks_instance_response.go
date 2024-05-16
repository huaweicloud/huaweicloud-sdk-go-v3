package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteStarrocksInstanceResponse Response Object
type DeleteStarrocksInstanceResponse struct {

	// 工作流ID。
	WorkflowId     *string `json:"workflow_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteStarrocksInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteStarrocksInstanceResponse struct{}"
	}

	return strings.Join([]string{"DeleteStarrocksInstanceResponse", string(data)}, " ")
}
