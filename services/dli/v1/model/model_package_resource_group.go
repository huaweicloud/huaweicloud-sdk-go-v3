package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PackageResourceGroup struct {
	GroupName *string `json:"group_name,omitempty"`

	Status *string `json:"status,omitempty"`

	Resources *[]string `json:"resources,omitempty"`

	Details *[]PackageResource `json:"details,omitempty"`

	CreateTime *int64 `json:"create_time,omitempty"`

	UpdateTime *int64 `json:"update_time,omitempty"`

	Owner *string `json:"owner,omitempty"`

	IsAsync *bool `json:"is_async,omitempty"`
}

func (o PackageResourceGroup) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PackageResourceGroup struct{}"
	}

	return strings.Join([]string{"PackageResourceGroup", string(data)}, " ")
}
