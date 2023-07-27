package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunCreateInstanceRequest Request Object
type RunCreateInstanceRequest struct {

	// 用户的project_id  登陆华为云 -> 用户中心 -> 我的凭证 -> api凭证 即可查看对应区域的项目ID。
	ProjectId string `json:"project_id"`

	Body *CreateInstanceReq `json:"body,omitempty"`
}

func (o RunCreateInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunCreateInstanceRequest struct{}"
	}

	return strings.Join([]string{"RunCreateInstanceRequest", string(data)}, " ")
}
