package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateRepositoryLabelResponse Response Object
type CreateRepositoryLabelResponse struct {

	// **参数解释：** 标签ID。 **取值范围：** 1-2147483647
	Id *int32 `json:"id,omitempty"`

	// **参数解释：** 标签名。 **取值范围：** 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释：** 标签颜色，以6位十六进制表示法给出，带有前导“#”符号（例如，#FFAABB）。 **取值范围：** 正则`^#([a-fA-F0-9]{6}|[a-fA-F0-9]{3})$`
	Color *string `json:"color,omitempty"`

	// **参数解释：** 描述。 **取值范围：** 不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释：** 字体颜色，以6位十六进制表示法给出，带有前导“#”符号（例如，#FFAABB）。 **取值范围：** 正则`^#([a-fA-F0-9]{6}|[a-fA-F0-9]{3})$`
	TextColor *string `json:"text_color,omitempty"`

	// **参数解释：** 失效时间。 **取值范围：** 不涉及。
	ExpiresAt *string `json:"expires_at,omitempty"`

	// **参数解释：** 是否失效。 **取值范围：** 不涉及。
	IsExpired *bool `json:"is_expired,omitempty"`

	// **参数解释：** 关联开启状态MR的数量。 **约束限制：** MR仓库返回此字段。 **取值范围：** 0-2147483647
	OpenMergeRequestsCount *int32 `json:"open_merge_requests_count,omitempty"`

	// **参数解释：**  关联开启状态CR的数量。  **约束限制：**  CR仓库返回此字段。  **取值范围：**  0-2147483647
	OpenChangeRequestCount *int32 `json:"open_change_request_count,omitempty"`

	// **参数解释：** 优先级 **取值范围：** 不涉及
	Priority *int32 `json:"priority,omitempty"`

	// **参数解释：** 是否为仓库标签 **取值范围：** 不涉及
	IsRepositoryLabel *bool `json:"is_repository_label,omitempty"`
	HttpStatusCode    int   `json:"-"`
}

func (o CreateRepositoryLabelResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRepositoryLabelResponse struct{}"
	}

	return strings.Join([]string{"CreateRepositoryLabelResponse", string(data)}, " ")
}
