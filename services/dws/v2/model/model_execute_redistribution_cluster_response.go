package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ExecuteRedistributionClusterResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ExecuteRedistributionClusterResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteRedistributionClusterResponse struct{}"
	}

	return strings.Join([]string{"ExecuteRedistributionClusterResponse", string(data)}, " ")
}
