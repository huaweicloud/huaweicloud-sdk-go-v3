package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateDeviceGroupRequestBody struct {

	// 分组名称，支持中文，英文大小写，数字，下划线和中划线,长度2-64
	Name string `json:"name"`

	// 分组描述，长度0-200
	Description *string `json:"description,omitempty"`
}

func (o UpdateDeviceGroupRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDeviceGroupRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateDeviceGroupRequestBody", string(data)}, " ")
}
