package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type OpenTag struct {
	Name string `json:"name"`

	Description string `json:"description"`

	TagId string `json:"tag_id"`

	CreateTime *int64 `json:"create_time,omitempty"`

	UpdateTime *int64 `json:"update_time,omitempty"`

	CreateUser *string `json:"create_user,omitempty"`

	DomainId *string `json:"domain_id,omitempty"`

	InstanceId *string `json:"instance_id,omitempty"`

	ProjectId *string `json:"project_id,omitempty"`

	CreateUserId *string `json:"create_user_id,omitempty"`
}

func (o OpenTag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OpenTag struct{}"
	}

	return strings.Join([]string{"OpenTag", string(data)}, " ")
}
