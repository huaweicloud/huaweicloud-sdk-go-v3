package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListJobResponse Response Object
type ListJobResponse struct {
	Data           *GetCreateFirewallJobResponseData `json:"data,omitempty"`
	HttpStatusCode int                               `json:"-"`
}

func (o ListJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListJobResponse struct{}"
	}

	return strings.Join([]string{"ListJobResponse", string(data)}, " ")
}
