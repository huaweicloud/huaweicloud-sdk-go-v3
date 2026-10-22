package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ObsConfig **参数解释**：OBS存储配置。 **约束限制**：不涉及。
type ObsConfig struct {

	// **参数解释**：OBS桶名称。该桶的类型必须是并行文件系统。 **约束限制**：不涉及。 **取值范围**：长度为3~63个字符，支持小写字母、数字、中划线和句点的组合。 **默认取值**：不涉及。
	BucketName *string `json:"bucket_name,omitempty"`

	// **参数解释**：委托名称。云服务将使用该委托操作OBS桶。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	AgencyName *string `json:"agency_name,omitempty"`
}

func (o ObsConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ObsConfig struct{}"
	}

	return strings.Join([]string{"ObsConfig", string(data)}, " ")
}
