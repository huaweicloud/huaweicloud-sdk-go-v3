package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateDeviceGroupRequestBody struct {
	// 父分组ID，自动向下取整

	ParentId int32 `json:"parent_id"`
	// 分组名称，支持中文，英文大小写，数字，下划线和中划线,长度2-64

	Name string `json:"name"`
	// 分组描述

	Description *string `json:"description,omitempty"`
	// 分组归属应用ID

	AppId string `json:"app_id"`
}

func (o CreateDeviceGroupRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDeviceGroupRequestBody struct{}"
	}

	return strings.Join([]string{"CreateDeviceGroupRequestBody", string(data)}, " ")
}
