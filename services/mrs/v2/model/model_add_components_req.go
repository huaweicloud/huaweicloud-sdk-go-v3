package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddComponentsReq 添加组件请求
type AddComponentsReq struct {

	// 组件模型详情
	ComponentsInstallMode []ComponentInstallMode `json:"components_install_mode"`
}

func (o AddComponentsReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddComponentsReq struct{}"
	}

	return strings.Join([]string{"AddComponentsReq", string(data)}, " ")
}
