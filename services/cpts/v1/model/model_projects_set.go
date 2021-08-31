package model

import (
	"encoding/json"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"

	"strings"
)

type ProjectsSet struct {
	// 创建时间

	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`
	// 更新时间

	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`
	// 描述

	Description *string `json:"description,omitempty"`
	// 工程id

	Id *int32 `json:"id,omitempty"`
	// 工程名字

	Name *string `json:"name,omitempty"`
	// 工程状态

	Status *int32 `json:"status,omitempty"`
	// 工程所属组

	Group *string `json:"group,omitempty"`
}

func (o ProjectsSet) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ProjectsSet struct{}"
	}

	return strings.Join([]string{"ProjectsSet", string(data)}, " ")
}
