package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateLineGroupsRequest Request Object
type UpdateLineGroupsRequest struct {

	// 线路分组ID
	LinegroupId string `json:"linegroup_id"`

	Body *UpdateLineGroupsRequestBody `json:"body,omitempty"`
}

func (o UpdateLineGroupsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLineGroupsRequest struct{}"
	}

	return strings.Join([]string{"UpdateLineGroupsRequest", string(data)}, " ")
}
