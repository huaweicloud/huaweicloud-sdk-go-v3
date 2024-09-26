package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type GlobalConnectionBandwidthLineLevel struct {

	// 实例ID。
	Id string `json:"id"`

	// 实例描述。不支持 <>。
	Description *string `json:"description,omitempty"`

	// 实例创建时间。UTC时间格式，yyyy-MM-ddTHH:mm:ss。
	CreatedAt *sdktime.SdkTime `json:"created_at"`

	// 实例更新时间。UTC时间格式，yyyy-MM-ddTHH:mm:ss。
	UpdatedAt *sdktime.SdkTime `json:"updated_at"`

	// 功能说明：线路分级本端接入点。
	LocalArea *string `json:"local_area,omitempty"`

	// 功能描述：线路分级远端接入点。
	RemoteArea *string `json:"remote_area,omitempty"`

	// 支持的铂金、金、银分级。
	Levels *[]string `json:"levels,omitempty"`
}

func (o GlobalConnectionBandwidthLineLevel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GlobalConnectionBandwidthLineLevel struct{}"
	}

	return strings.Join([]string{"GlobalConnectionBandwidthLineLevel", string(data)}, " ")
}
