package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunSearchPictureRequest Request Object
type RunSearchPictureRequest struct {

	// 实例名称。
	InstanceName string `json:"instance_name"`

	// 用户的project_id 登陆华为云 -> 用户中心 -> 我的凭证 -> api凭证 即可查看对应区域的项目ID。
	ProjectId string `json:"project_id"`

	Body *SearchPictureReq `json:"body,omitempty"`
}

func (o RunSearchPictureRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunSearchPictureRequest struct{}"
	}

	return strings.Join([]string{"RunSearchPictureRequest", string(data)}, " ")
}
