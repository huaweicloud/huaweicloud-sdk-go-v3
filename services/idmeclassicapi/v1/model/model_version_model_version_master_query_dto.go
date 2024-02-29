package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VersionModelVersionMasterQueryDto struct {

	// 是否加密。 - true：加密。 - false：不加密。
	Decrypt *bool `json:"decrypt,omitempty"`

	// 迭代版本。
	Iteration *int32 `json:"iteration,omitempty"`

	// 主对象ID。
	MasterId *int64 `json:"masterId,omitempty"`

	// 版本号。
	Version *string `json:"version,omitempty"`
}

func (o VersionModelVersionMasterQueryDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VersionModelVersionMasterQueryDto struct{}"
	}

	return strings.Join([]string{"VersionModelVersionMasterQueryDto", string(data)}, " ")
}
