package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VersionModelMasterIdsDto struct {

	// 父模型ID。
	MasterId *string `json:"masterId,omitempty"`

	// 版本对象。
	Version *string `json:"version,omitempty"`
}

func (o VersionModelMasterIdsDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VersionModelMasterIdsDto struct{}"
	}

	return strings.Join([]string{"VersionModelMasterIdsDto", string(data)}, " ")
}
