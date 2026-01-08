package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateHostsRequest Request Object
type UpdateHostsRequest struct {
	Body *UpdateHostsRequestBody `json:"body,omitempty"`
}

func (o UpdateHostsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateHostsRequest struct{}"
	}

	return strings.Join([]string{"UpdateHostsRequest", string(data)}, " ")
}
