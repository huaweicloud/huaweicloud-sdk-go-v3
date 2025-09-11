package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpgradeInstancesVersionRequest Request Object
type UpgradeInstancesVersionRequest struct {

	// **参数解释**: 语言。 **约束限制**: 不涉及。 **取值范围**:   - zh-cn   - en-us  **默认取值**: en-us
	XLanguage *string `json:"X-Language,omitempty"`

	Body *GaussDbUpgradeInstancesVersionRequest `json:"body,omitempty"`
}

func (o UpgradeInstancesVersionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpgradeInstancesVersionRequest struct{}"
	}

	return strings.Join([]string{"UpgradeInstancesVersionRequest", string(data)}, " ")
}
