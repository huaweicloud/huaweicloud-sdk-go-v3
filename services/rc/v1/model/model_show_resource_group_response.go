package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowResourceGroupResponse Response Object
type ShowResourceGroupResponse struct {

	// uuid
	GroupId *string `json:"group_id,omitempty"`

	DomainId *string `json:"domain_id,omitempty"`

	GroupName *string `json:"group_name,omitempty"`

	Description *string `json:"description,omitempty"`

	CreateTime *string `json:"create_time,omitempty"`

	UpdateTime     *string `json:"update_time,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowResourceGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowResourceGroupResponse struct{}"
	}

	return strings.Join([]string{"ShowResourceGroupResponse", string(data)}, " ")
}
