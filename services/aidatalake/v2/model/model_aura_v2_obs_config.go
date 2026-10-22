package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AuraV2ObsConfig **参数解释**：OBS存储配置。
type AuraV2ObsConfig struct {

	// **参数解释**：OBS桶名称。 **约束限制**：不涉及。 **取值范围**：长度为3~63个字符，支持小写字母、数字、中划线和句点的组合。 **默认取值**：不涉及。
	BucketName string `json:"bucket_name"`
}

func (o AuraV2ObsConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AuraV2ObsConfig struct{}"
	}

	return strings.Join([]string{"AuraV2ObsConfig", string(data)}, " ")
}
