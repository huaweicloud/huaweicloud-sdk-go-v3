package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatedAt 实例更新时间。
type UpdatedAt struct {

	// 实例更新时间。UTC时间格式，yyyy-MM-ddTHH:mm:ss。
	UpdatedAt *sdktime.SdkTime `json:"updated_at"`
}

func (o UpdatedAt) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatedAt struct{}"
	}

	return strings.Join([]string{"UpdatedAt", string(data)}, " ")
}
