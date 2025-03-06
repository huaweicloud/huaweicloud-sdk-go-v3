package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ModifyApplicationRequestBody struct {

	// 应用ID。
	Id string `json:"id"`

	// 应用的中文描述。
	Description string `json:"description"`

	// 应用的英文描述。
	DescriptionEn string `json:"description_en"`

	// app权限控制。 - NONE：关闭权限校验 - ALL：开启所有校验
	PermissionControl *string `json:"permission_control,omitempty"`

	// 应用责任人。
	AppUserList []AppUserList `json:"app_user_list"`
}

func (o ModifyApplicationRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyApplicationRequestBody struct{}"
	}

	return strings.Join([]string{"ModifyApplicationRequestBody", string(data)}, " ")
}
