package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFeaturesResponse Response Object
type ListFeaturesResponse struct {

	// **参数解释**：实例创建的时间，UTC毫秒。 **取值范围**：不涉及。
	CreateAt *int64 `json:"create_at,omitempty"`

	// **参数解释**：特性开关。 **取值范围**：布尔类型： - true：开启。 - false：未开启。
	Enabled *bool `json:"enabled,omitempty"`

	// **参数解释**：特性ID。 **取值范围**：不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**：特性名称。 **取值范围**：枚举类型，取值如下： - NOTEBOOK：用户显式创建的Notebook实例。
	Name *string `json:"name,omitempty"`

	// **参数解释**：特性配额。 **取值范围**：不涉及。
	Quota *int32 `json:"quota,omitempty"`

	// **参数解释**：特性已使用额度。 **取值范围**：不涉及。
	Used *int32 `json:"used,omitempty"`

	// **参数解释**：实例最后更新的时间，UTC毫秒。 **取值范围**：不涉及。
	UpdateAt *int64 `json:"update_at,omitempty"`

	// **参数解释**：用户ID。 **取值范围**：不涉及。
	UserId         *string `json:"user_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListFeaturesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFeaturesResponse struct{}"
	}

	return strings.Join([]string{"ListFeaturesResponse", string(data)}, " ")
}
