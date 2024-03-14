package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VersionModelVersionMasterQueryDto struct {

	// 是否加密。 - true：加密。 - false：不加密。
	Decrypt *bool `json:"decrypt,omitempty"`

	// 迭代版本。如果此参数为空，则返回M-V模型实例的最新版本信息。
	Iteration *int32 `json:"iteration,omitempty"`

	// 主对象ID。
	MasterId string `json:"masterId"`

	// 版本号。
	Version string `json:"version"`
}

func (o VersionModelVersionMasterQueryDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VersionModelVersionMasterQueryDto struct{}"
	}

	return strings.Join([]string{"VersionModelVersionMasterQueryDto", string(data)}, " ")
}
