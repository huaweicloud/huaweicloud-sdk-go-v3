package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSubnetIdsResponse Response Object
type UpdateSubnetIdsResponse struct {

	// 任务id
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateSubnetIdsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSubnetIdsResponse struct{}"
	}

	return strings.Join([]string{"UpdateSubnetIdsResponse", string(data)}, " ")
}
