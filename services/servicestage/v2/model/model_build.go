package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 构建工程。
type Build struct {
	Parameters *BuildParameters `json:"parameters,omitempty"`
}

func (o Build) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Build struct{}"
	}

	return strings.Join([]string{"Build", string(data)}, " ")
}
