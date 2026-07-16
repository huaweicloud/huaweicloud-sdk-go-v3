package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AlgorithmResponseMetadata 算法的元数据，描述算法基本信息。
type AlgorithmResponseMetadata struct {

	// **参数解释**：算法id，创建算法时无需填写。 **取值范围**：不涉及。
	Id *string `json:"id,omitempty"`

	// 算法名称。限制为1-64位只含数字、字母、下划线和中划线的名称。
	Name string `json:"name"`

	// 对算法的描述，默认为“NULL”，字符串的长度限制为[0, 256]。
	Description *string `json:"description,omitempty"`

	// 指定算法所处的工作空间，默认值为“0”。“0” 为默认的工作空间。
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// 指定算法所属的ai项目，默认值为\"default-ai-project\"。ai项目已下线，无需关注。
	AiProject *string `json:"ai_project,omitempty"`

	// 用户名称。
	UserName *string `json:"user_name,omitempty"`

	// 用户的domainID。
	DomainId *string `json:"domain_id,omitempty"`

	// 算法来源类型。
	Source *string `json:"source,omitempty"`

	// 算法api版本，标识新旧版。
	ApiVersion *string `json:"api_version,omitempty"`

	// **参数解释**：算法可用性。 **取值范围**： - true：可用 - false：不可用
	IsValid *bool `json:"is_valid,omitempty"`

	// 算法状态。
	State *string `json:"state,omitempty"`

	// 算法标签。
	Tags *[]map[string]string `json:"tags,omitempty"`

	// 算法属性列表。
	AttrList *[]string `json:"attr_list,omitempty"`

	// 算法版本数量，默认为0。
	VersionNum *int32 `json:"version_num,omitempty"`

	// 算法大小。
	Size *int32 `json:"size,omitempty"`

	// 算法创建时间戳。
	CreateTime *int64 `json:"create_time,omitempty"`

	// 算法更新时间戳。
	UpdateTime *int64 `json:"update_time,omitempty"`
}

func (o AlgorithmResponseMetadata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlgorithmResponseMetadata struct{}"
	}

	return strings.Join([]string{"AlgorithmResponseMetadata", string(data)}, " ")
}
