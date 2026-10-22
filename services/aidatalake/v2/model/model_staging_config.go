package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StagingConfig **参数解释**：临时存储配置。 **约束限制**：不涉及。
type StagingConfig struct {
	Obs *ObsConfig `json:"obs,omitempty"`
}

func (o StagingConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StagingConfig struct{}"
	}

	return strings.Join([]string{"StagingConfig", string(data)}, " ")
}
