package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeployFactoryPackagesRequestBody struct {

	// 发布包ID
	PackageIds []string `json:"package_ids"`

	// 发布后是否立即启动作业。取值范围为0和1，默认为1, 1：发布成功后立即启动作业 0：不立即启动
	StartupMode *int32 `json:"startup_mode,omitempty"`
}

func (o DeployFactoryPackagesRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeployFactoryPackagesRequestBody struct{}"
	}

	return strings.Join([]string{"DeployFactoryPackagesRequestBody", string(data)}, " ")
}
