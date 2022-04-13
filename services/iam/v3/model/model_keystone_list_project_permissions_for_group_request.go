package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type KeystoneListProjectPermissionsForGroupRequest struct {
	// 项目ID，获取方式请参见：[获取项目名称、项目ID](https://support.huaweicloud.com/api-iam/iam_17_0002.html)。

	ProjectId string `json:"project_id"`
	// 用户组ID，获取方式请参见：[获取用户组ID](https://support.huaweicloud.com/api-iam/iam_17_0002.html)。

	GroupId string `json:"group_id"`
}

func (o KeystoneListProjectPermissionsForGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeystoneListProjectPermissionsForGroupRequest struct{}"
	}

	return strings.Join([]string{"KeystoneListProjectPermissionsForGroupRequest", string(data)}, " ")
}
