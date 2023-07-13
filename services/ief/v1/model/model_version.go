package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Version 应用模板版本
type Version struct {
	Version *VersionDetail `json:"version"`
}

func (o Version) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Version struct{}"
	}

	return strings.Join([]string{"Version", string(data)}, " ")
}
