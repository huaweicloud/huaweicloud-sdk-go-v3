package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowGroupResponse Response Object
type ShowGroupResponse struct {
	Group          *DescribeGroupsRespGroup `json:"group,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o ShowGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowGroupResponse struct{}"
	}

	return strings.Join([]string{"ShowGroupResponse", string(data)}, " ")
}
