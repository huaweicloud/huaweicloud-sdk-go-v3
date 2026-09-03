package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CurrentPackageInfo 当前生效技能包信息（用于批量查询响应）。
type CurrentPackageInfo struct {

	// 技能包id。
	Id *string `json:"id,omitempty"`

	// 版本号。
	Version *string `json:"version,omitempty"`

	// 版本修订号。
	Revision *int32 `json:"revision,omitempty"`

	PackageStatus *PackageStatusEnum `json:"package_status,omitempty"`

	// 已部署的区域标识列表。
	Regions *[]string `json:"regions,omitempty"`
}

func (o CurrentPackageInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CurrentPackageInfo struct{}"
	}

	return strings.Join([]string{"CurrentPackageInfo", string(data)}, " ")
}
