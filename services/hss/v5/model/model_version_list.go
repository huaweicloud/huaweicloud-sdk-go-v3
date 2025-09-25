package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VersionList struct {

	// 版本号
	ReleaseVersion *string `json:"release_version,omitempty"`

	// 版本说明
	ReleaseNote *string `json:"release_note,omitempty"`

	// 更新时间，毫秒
	UpdateTime *int64 `json:"update_time,omitempty"`
}

func (o VersionList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VersionList struct{}"
	}

	return strings.Join([]string{"VersionList", string(data)}, " ")
}
