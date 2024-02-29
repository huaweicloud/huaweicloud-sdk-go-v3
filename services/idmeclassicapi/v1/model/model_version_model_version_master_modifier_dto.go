package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VersionModelVersionMasterModifierDto struct {

	// 父模型ID。
	MasterId *int64 `json:"masterId,omitempty"`

	// 修改人。
	Modifier *string `json:"modifier,omitempty"`

	// 版本对象。
	Version *string `json:"version,omitempty"`
}

func (o VersionModelVersionMasterModifierDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VersionModelVersionMasterModifierDto struct{}"
	}

	return strings.Join([]string{"VersionModelVersionMasterModifierDto", string(data)}, " ")
}
