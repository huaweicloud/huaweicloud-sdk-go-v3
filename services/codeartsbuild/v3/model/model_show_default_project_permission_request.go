package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDefaultProjectPermissionRequest Request Object
type ShowDefaultProjectPermissionRequest struct {

	// CodeArts项目ID，32位数字、小写字母组合。
	ProjectId string `json:"project_id"`

	// 构建的任务ID； 编辑构建任务时，浏览器URL末尾的32位数字、字母组合的字符串。
	JobId string `json:"job_id"`
}

func (o ShowDefaultProjectPermissionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDefaultProjectPermissionRequest struct{}"
	}

	return strings.Join([]string{"ShowDefaultProjectPermissionRequest", string(data)}, " ")
}
