package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type LtsConfig struct {

	// AccessLog日志组ID
	LogGroupID *string `json:"logGroupID,omitempty"`

	// AccessLog输出日志流ID
	LogStreamID *string `json:"logStreamID,omitempty"`
}

func (o LtsConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LtsConfig struct{}"
	}

	return strings.Join([]string{"LtsConfig", string(data)}, " ")
}
