package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateLineGroupsRequest struct {

	// 待更新的线路分组ID。
	LinegroupId string `json:"linegroup_id"`
}

func (o UpdateLineGroupsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLineGroupsRequest struct{}"
	}

	return strings.Join([]string{"UpdateLineGroupsRequest", string(data)}, " ")
}
