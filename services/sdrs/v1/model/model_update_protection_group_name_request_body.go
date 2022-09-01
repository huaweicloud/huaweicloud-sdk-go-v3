package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 更新保护组名称请求体
type UpdateProtectionGroupNameRequestBody struct {
	ServerGroup *UpdateProtectionGroupNameRequestParams `json:"server_group" xml:"server_group"`
}

func (o UpdateProtectionGroupNameRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateProtectionGroupNameRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateProtectionGroupNameRequestBody", string(data)}, " ")
}
