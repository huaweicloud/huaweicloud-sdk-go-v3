package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateRepoDomainsRequestBody struct {
	// 当前只支持read权限

	Permit string `json:"permit"`
	// 截止时间，UTC时间格式。永久有效为forever

	Deadline string `json:"deadline"`
	// 描述。默认值为空字符串

	Description *string `json:"description,omitempty"`
}

func (o UpdateRepoDomainsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRepoDomainsRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateRepoDomainsRequestBody", string(data)}, " ")
}
