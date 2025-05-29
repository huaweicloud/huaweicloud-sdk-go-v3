package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowProjectPermissionRequest Request Object
type ShowProjectPermissionRequest struct {

	// CodeArts项目ID，32位数字、小写字母组合。
	ProjectId string `json:"project_id"`
}

func (o ShowProjectPermissionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowProjectPermissionRequest struct{}"
	}

	return strings.Join([]string{"ShowProjectPermissionRequest", string(data)}, " ")
}
