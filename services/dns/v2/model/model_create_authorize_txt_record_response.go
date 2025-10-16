package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAuthorizeTxtRecordResponse Response Object
type CreateAuthorizeTxtRecordResponse struct {

	// **参数解释：** 授权请求ID。 **取值范围：** 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释：** 待创建的子域名。 **取值范围：** 不涉及。
	ZoneName *string `json:"zone_name,omitempty"`

	Record *RecordInfo `json:"record,omitempty"`

	// **参数解释：** 授权状态。 **取值范围：** - CREATED：已创建 - VERIFIED：已验证
	Status *string `json:"status,omitempty"`

	// **参数解释：** 子域名所属的二级域名。 **取值范围：** 不涉及。
	SecondLevelZoneName *string `json:"second_level_zone_name,omitempty"`

	// **参数解释：** 创建时间。 格式：yyyy-MM-dd'T'HH:mm:ss.SSS。 **取值范围：** 不涉及。
	CreatedAt *string `json:"created_at,omitempty"`

	// **参数解释：** 更新时间。 格式：yyyy-MM-dd'T'HH:mm:ss.SSS。 **取值范围：** 不涉及。
	UpdatedAt      *string `json:"updated_at,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateAuthorizeTxtRecordResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAuthorizeTxtRecordResponse struct{}"
	}

	return strings.Join([]string{"CreateAuthorizeTxtRecordResponse", string(data)}, " ")
}
