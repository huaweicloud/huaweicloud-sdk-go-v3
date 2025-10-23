package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RepositoryTemplateDto **参数解释：** 模板仓列表
type RepositoryTemplateDto struct {

	// **参数解释：** 仓库Id。
	RepositoryId *int32 `json:"repository_id,omitempty"`

	// **参数解释：** 模板仓标题。 **取值范围：** 字符串长度1-50。
	Name *string `json:"name,omitempty"`

	// **参数解释：** 是否是系统模板。 **取值范围：** - true，系统模板。 - false，个人模板。
	System *bool `json:"system,omitempty"`

	// **参数解释：** 标签列表。 **取值范围：** 不涉及。
	Tags *[]string `json:"tags,omitempty"`

	// **参数解释：** 仓库描述信息。 **取值范围：** 字符串长度0-2000。
	Description *string `json:"description,omitempty"`

	// **参数解释：** 编程语言。 **取值范围：** 字符串长度0-32。
	Language *string `json:"language,omitempty"`

	// **参数解释：** 模板仓的仓库名称。 **取值范围：** 字符串长度0-255。
	RepositoryName *string `json:"repository_name,omitempty"`

	// **参数解释：** 模板仓的简介。 **取值范围：** 字符串长度0-32。
	BriefIntroduction *string `json:"brief_introduction,omitempty"`

	// **参数解释：** 创建时间。
	CreatedAt *string `json:"created_at,omitempty"`

	// **参数解释：** 模板仓被使用的次数。
	UsedTimes *int32 `json:"used_times,omitempty"`

	// **参数解释：** 模板仓被点赞的次数。
	LikedTimes *int32 `json:"liked_times,omitempty"`

	// **参数解释：** 作者。 **取值范围：** 字符串长度0-128。
	CreatorName *string `json:"creator_name,omitempty"`

	// **参数解释：** 代码仓https协议的git地址。 **取值范围：** 字符串最大长度512。
	HttpsUrl *string `json:"https_url,omitempty"`
}

func (o RepositoryTemplateDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RepositoryTemplateDto struct{}"
	}

	return strings.Join([]string{"RepositoryTemplateDto", string(data)}, " ")
}
