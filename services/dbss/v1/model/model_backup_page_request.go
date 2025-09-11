package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BackupPageRequest struct {

	// 结束时间
	EndTime *sdktime.SdkTime `json:"end_time,omitempty"`

	// 当前页码
	Page int32 `json:"page"`

	// 分页大小
	Size int32 `json:"size"`

	// 开始时间,yyyy-MM-dd HH:mm:ss
	StartTime *sdktime.SdkTime `json:"start_time,omitempty"`
}

func (o BackupPageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BackupPageRequest struct{}"
	}

	return strings.Join([]string{"BackupPageRequest", string(data)}, " ")
}
