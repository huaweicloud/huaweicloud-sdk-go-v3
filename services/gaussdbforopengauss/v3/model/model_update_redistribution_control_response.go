package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateRedistributionControlResponse Response Object
type UpdateRedistributionControlResponse struct {

	// 任务流id
	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateRedistributionControlResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRedistributionControlResponse struct{}"
	}

	return strings.Join([]string{"UpdateRedistributionControlResponse", string(data)}, " ")
}
