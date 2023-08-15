package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunAddPictureRequest Request Object
type RunAddPictureRequest struct {

	// 用户的project_id 登陆华为云 -> 用户中心 -> 我的凭证 -> api凭证 即可查看对应区域的项目ID。
	ProjectId string `json:"project_id"`

	// 实例名称。
	InstanceName string `json:"instance_name"`

	Body *AddPictureRequestReq `json:"body,omitempty"`
}

func (o RunAddPictureRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunAddPictureRequest struct{}"
	}

	return strings.Join([]string{"RunAddPictureRequest", string(data)}, " ")
}
