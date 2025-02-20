package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ApplicationInfo struct {
	Id *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`

	Code *string `json:"code,omitempty"`

	Description *string `json:"description,omitempty"`

	DomainId *string `json:"domain_id,omitempty"`

	ParentId *string `json:"parent_id,omitempty"`

	Path *string `json:"path,omitempty"`

	CreateTime *string `json:"create_time,omitempty"`

	UpdateTime *string `json:"update_time,omitempty"`
}

func (o ApplicationInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplicationInfo struct{}"
	}

	return strings.Join([]string{"ApplicationInfo", string(data)}, " ")
}
