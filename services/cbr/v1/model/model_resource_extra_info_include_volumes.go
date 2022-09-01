package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type ResourceExtraInfoIncludeVolumes struct {

	// 卷ID，仅支持uuid
	Id string `json:"id" xml:"id"`

	// 操作系统类型
	OsVersion *string `json:"os_version,omitempty" xml:"os_version"`
}

func (o ResourceExtraInfoIncludeVolumes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceExtraInfoIncludeVolumes struct{}"
	}

	return strings.Join([]string{"ResourceExtraInfoIncludeVolumes", string(data)}, " ")
}
