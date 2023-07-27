package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunModifyPictureRequest Request Object
type RunModifyPictureRequest struct {

	// 实例名称。
	InstanceName string `json:"instance_name"`

	// 用户的project_id 登陆华为云 -> 用户中心 -> 我的凭证 -> api凭证 即可查看对应区域的项目ID。
	ProjectId string `json:"project_id"`

	Body *RunModifyPictureReq `json:"body,omitempty"`
}

func (o RunModifyPictureRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunModifyPictureRequest struct{}"
	}

	return strings.Join([]string{"RunModifyPictureRequest", string(data)}, " ")
}
