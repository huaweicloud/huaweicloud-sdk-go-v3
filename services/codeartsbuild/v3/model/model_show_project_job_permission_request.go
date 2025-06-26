package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowProjectJobPermissionRequest Request Object
type ShowProjectJobPermissionRequest struct {

	// CodeArts项目ID，32位数字、小写字母组合。
	ProjectId string `json:"project_id"`

	// 构建的任务ID； 编辑构建任务时，浏览器URL末尾的32位数字、字母组合的字符串。
	JobId string `json:"job_id"`
}

func (o ShowProjectJobPermissionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowProjectJobPermissionRequest struct{}"
	}

	return strings.Join([]string{"ShowProjectJobPermissionRequest", string(data)}, " ")
}
