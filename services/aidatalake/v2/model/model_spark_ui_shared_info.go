package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SparkUiSharedInfo Spark UI 共享链接及其到期时刻。
type SparkUiSharedInfo struct {

	// **参数解释**：共享链接凭证。 **约束限制**：不涉及。 **取值范围**：共享URL格式，长度为1~1024个字符。 **默认取值**：不涉及。
	SharedUrl string `json:"shared_url"`

	// **参数解释**：过期时间时间戳，共享链接使用到该时间之后不可用，毫秒。 **约束限制**：毫秒时间戳。 **取值范围**：不涉及。 **默认取值**：不涉及。
	ExpiredAt int64 `json:"expired_at"`
}

func (o SparkUiSharedInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SparkUiSharedInfo struct{}"
	}

	return strings.Join([]string{"SparkUiSharedInfo", string(data)}, " ")
}
