package model

import (
	"encoding/json"

	"strings"
)

//
type ResourceExtraInfoIncludeVolumes struct {
	// 卷ID，仅支持uuid

	Id string `json:"id"`
	// 操作系统类型

	OsVersion *string `json:"os_version,omitempty"`
}

func (o ResourceExtraInfoIncludeVolumes) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ResourceExtraInfoIncludeVolumes struct{}"
	}

	return strings.Join([]string{"ResourceExtraInfoIncludeVolumes", string(data)}, " ")
}
