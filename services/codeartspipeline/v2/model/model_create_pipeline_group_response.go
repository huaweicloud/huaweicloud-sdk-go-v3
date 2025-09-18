package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePipelineGroupResponse Response Object
type CreatePipelineGroupResponse struct {

	// **参数解释**： 分组ID。 **取值范围**： 32位字符，由数字和字母组成。
	Id *string `json:"id,omitempty"`

	// **参数解释**： 租户ID。 **取值范围**： 32位字符，由数字和字母组成。
	DomainId *string `json:"domain_id,omitempty"`

	// **参数解释**： 项目ID。 **取值范围**： 32位字符，由数字和字母组成。
	ProjectId *string `json:"project_id,omitempty"`

	// **参数解释**： 分组名。 **取值范围**： 32位字符，由数字和字母组成。
	Name *string `json:"name,omitempty"`

	// **参数解释**： 父分组ID。 **取值范围**： 32位字符，由数字和字母组成。
	ParentId *string `json:"parent_id,omitempty"`

	// **参数解释**： 分组路径ID。例如id1.id2.id3 代表分组id3的父分组为id2，分组id2的父分组为id1。 **取值范围**： 不涉及。
	PathId *string `json:"path_id,omitempty"`

	// **参数解释**： 序号。 **取值范围**： 大于等于1。
	Ordinal *int32 `json:"ordinal,omitempty"`

	// **参数解释**： 创建用户ID。 **取值范围**： 32位字符，由数字和字母组成。
	Creator *string `json:"creator,omitempty"`

	// **参数解释**： 更新用户ID。 **取值范围**： 32位字符，由数字和字母组成。
	Updater *string `json:"updater,omitempty"`

	// **参数解释**： 创建时间。 **取值范围**： 不涉及。
	CreateTime *int64 `json:"create_time,omitempty"`

	// **参数解释**： 更新时间。 **取值范围**： 不涉及。
	UpdateTime *int64 `json:"update_time,omitempty"`

	// **参数解释**： 流水线分组详情。 **取值范围**： 不涉及。
	Children       *[]PipelineGroupVo `json:"children,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o CreatePipelineGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePipelineGroupResponse struct{}"
	}

	return strings.Join([]string{"CreatePipelineGroupResponse", string(data)}, " ")
}
