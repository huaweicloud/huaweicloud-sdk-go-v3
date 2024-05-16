package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateStarrocksInstanceResponse Response Object
type CreateStarrocksInstanceResponse struct {
	Instance *SrCreateInstanceRspInstance `json:"instance,omitempty"`

	// 工作流ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateStarrocksInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateStarrocksInstanceResponse struct{}"
	}

	return strings.Join([]string{"CreateStarrocksInstanceResponse", string(data)}, " ")
}
