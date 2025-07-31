package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdatePublicZoneInfo struct {

	// **参数解释：** 域名的描述信息。 **约束限制：** 不涉及。 **取值范围：** 长度不超过255个字符。 **默认取值：** 默认为空，表示维持原值。
	Description *string `json:"description,omitempty"`

	// **参数解释：** 管理该域名的管理员邮箱，用于生成该域名的SOA记录。   **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 默认为空，表示维持原值。
	Email *string `json:"email,omitempty"`

	// **参数解释：** 用于填写默认生成的SOA记录中有效缓存时间，以秒为单位。 **约束限制：** 不涉及。 **取值范围：** 1~2147483647。 **默认取值：** 默认为空，表示维持原值。
	Ttl *int32 `json:"ttl,omitempty"`
}

func (o UpdatePublicZoneInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePublicZoneInfo struct{}"
	}

	return strings.Join([]string{"UpdatePublicZoneInfo", string(data)}, " ")
}
