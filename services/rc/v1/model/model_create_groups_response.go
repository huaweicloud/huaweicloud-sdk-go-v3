package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateGroupsResponse struct {

	// uuid
	GroupId *string `json:"group_id,omitempty"`

	DomainId *string `json:"domain_id,omitempty"`

	GroupName *string `json:"group_name,omitempty"`

	Description *string `json:"description,omitempty"`

	Created *string `json:"created,omitempty"`

	Updated *string `json:"updated,omitempty"`
}

func (o CreateGroupsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateGroupsResponse struct{}"
	}

	return strings.Join([]string{"CreateGroupsResponse", string(data)}, " ")
}
