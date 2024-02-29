package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VersionModelVersionMasterDto struct {

	// 主对象ID。
	MasterId *int64 `json:"masterId,omitempty"`

	// 版本对象版本号。
	Version *string `json:"version,omitempty"`
}

func (o VersionModelVersionMasterDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VersionModelVersionMasterDto struct{}"
	}

	return strings.Join([]string{"VersionModelVersionMasterDto", string(data)}, " ")
}
