package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VersionModelVersionUpdateDto struct {

	// 唯一标识。
	Id string `json:"id"`

	// 迭代版本。
	Iteration *int32 `json:"iteration,omitempty"`

	// 版本号。
	Version string `json:"version"`
}

func (o VersionModelVersionUpdateDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VersionModelVersionUpdateDto struct{}"
	}

	return strings.Join([]string{"VersionModelVersionUpdateDto", string(data)}, " ")
}
