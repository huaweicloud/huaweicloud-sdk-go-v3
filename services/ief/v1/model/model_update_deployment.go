package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 更新部署请求Body参数
type UpdateDeployment struct {
	Deployment *UpdatePodDeployment `json:"deployment,omitempty"`
	// 应用部署描述修改，只修改描述不需要传入deployment参数

	Description *string `json:"description,omitempty"`
}

func (o UpdateDeployment) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDeployment struct{}"
	}

	return strings.Join([]string{"UpdateDeployment", string(data)}, " ")
}
