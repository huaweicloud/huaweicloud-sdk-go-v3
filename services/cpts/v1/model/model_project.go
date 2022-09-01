package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Project struct {

	// create_time
	CreateTime *string `json:"create_time,omitempty" xml:"create_time"`

	// description
	Description *string `json:"description,omitempty" xml:"description"`

	// group
	Group *string `json:"group,omitempty" xml:"group"`

	// id
	Id *int32 `json:"id,omitempty" xml:"id"`

	// name
	Name *string `json:"name,omitempty" xml:"name"`

	// source
	Source *int32 `json:"source,omitempty" xml:"source"`

	// update_time
	UpdateTime *string `json:"update_time,omitempty" xml:"update_time"`
}

func (o Project) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Project struct{}"
	}

	return strings.Join([]string{"Project", string(data)}, " ")
}
