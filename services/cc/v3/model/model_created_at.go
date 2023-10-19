package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatedAt 实例创建时间。
type CreatedAt struct {

	// 实例创建时间。UTC时间格式，yyyy-MM-ddTHH:mm:ss。
	CreatedAt *sdktime.SdkTime `json:"created_at"`
}

func (o CreatedAt) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatedAt struct{}"
	}

	return strings.Join([]string{"CreatedAt", string(data)}, " ")
}
