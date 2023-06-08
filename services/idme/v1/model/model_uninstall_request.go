package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UninstallRequest struct {

	// 运行服务的ID。
	EnvId string `json:"env_id"`

	// 待卸载的应用ID。
	AppId string `json:"app_id"`
}

func (o UninstallRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UninstallRequest struct{}"
	}

	return strings.Join([]string{"UninstallRequest", string(data)}, " ")
}
