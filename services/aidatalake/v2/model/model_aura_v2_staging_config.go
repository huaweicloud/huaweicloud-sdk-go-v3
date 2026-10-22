package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AuraV2StagingConfig **参数解释**：临时存储配置。
type AuraV2StagingConfig struct {
	Obs *AuraV2ObsConfig `json:"obs"`
}

func (o AuraV2StagingConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AuraV2StagingConfig struct{}"
	}

	return strings.Join([]string{"AuraV2StagingConfig", string(data)}, " ")
}
