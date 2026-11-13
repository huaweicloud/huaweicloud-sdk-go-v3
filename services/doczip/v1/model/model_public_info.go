package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PublicInfo struct {

	// **参数解释**： Key ID。 **约束限制**： 不涉及。 **取值范围**： [1-128]个字符。 **默认取值**： 不涉及。
	Kid *string `json:"kid,omitempty"`

	// **参数解释**： PEM 格式公钥。 **约束限制**： 不涉及。 **取值范围**： [1-2048]的字符串。 **默认取值**： 不涉及。
	PublicKey *string `json:"public_key,omitempty"`

	// **参数解释**： [创建时间。格式为YYYY-MM-DD HH:mm:ss。例如：2023-11-24 17:35:25。](tag:his) [创建时间。UTC时间。格式为YYYY-MM-DD HH:mm:ss。例如：2023-11-24 17:35:25。](tag:hws) **取值范围**： 不涉及。
	CreatedAt *sdktime.SdkTime `json:"created_at,omitempty"`

	// **参数解释**： [创建时间。格式为YYYY-MM-DD HH:mm:ss。例如：2023-11-24 17:35:25。](tag:his) [创建时间。UTC时间。格式为YYYY-MM-DD HH:mm:ss。例如：2023-11-24 17:35:25。](tag:hws) **取值范围**： 不涉及。
	UpdatedAt *sdktime.SdkTime `json:"updated_at,omitempty"`
}

func (o PublicInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublicInfo struct{}"
	}

	return strings.Join([]string{"PublicInfo", string(data)}, " ")
}
