package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ReleaseAppInfo 应用发布信息
type ReleaseAppInfo struct {

	// **参数解释**： 应用ID。 **取值范围**： [1-256]个字符。
	Id *string `json:"id,omitempty"`

	// **参数解释**： 应用名称。 **取值范围**： [1-200]个字符。
	Name *string `json:"name,omitempty"`

	// **参数解释**： 应用状态。 **取值范围**： - 0：待发布。 - 1：已发布，生效状态。 - 2：已发布，失效状态。
	ReleaseState *string `json:"release_state,omitempty"`

	// **参数解释**： 发布说明。 **取值范围**： [1-500]字符。
	ReleaseNote *string `json:"release_note,omitempty"`

	// **参数解释**： 创建用户ID。 **取值范围**： [1-50]个字符。
	CreateUserId *string `json:"create_user_id,omitempty"`

	// **参数解释**： [创建时间。格式为YYYY-MM-DD HH:mm:ss。例如：2023-11-24 17:35:25。](tag:his) [创建时间。UTC时间。格式为YYYY-MM-DD HH:mm:ss。例如：2023-11-24 17:35:25。](tag:hws) **取值范围**： 不涉及。
	CreatedAt *string `json:"created_at,omitempty"`

	// **参数解释**： 创建用户ID。 **取值范围**： [1-50]个字符。
	UpdateUserId *string `json:"update_user_id,omitempty"`

	// **参数解释**： [更新时间。格式为YYYY-MM-DD HH:mm:ss。例如：2023-11-24 17:35:25。](tag:his) [更新时间。UTC时间。格式为YYYY-MM-DD HH:mm:ss。例如：2023-11-24 17:35:25。](tag:hws) **取值范围**： 不涉及。
	UpdatedAt *string `json:"updated_at,omitempty"`

	// **参数解释**： 版本号。 **取值范围**： [1-50]个字符。
	Version *string `json:"version,omitempty"`
}

func (o ReleaseAppInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReleaseAppInfo struct{}"
	}

	return strings.Join([]string{"ReleaseAppInfo", string(data)}, " ")
}
