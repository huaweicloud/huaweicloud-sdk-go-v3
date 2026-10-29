package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CriterionSetListV2Body struct {

	// **参数解释：** 规则集ID。 **取值范围：** 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释：** 规则集中文名称。 **取值范围：** 任意中英文字符，数字，点，连字符和下划线组成的字符串，长度为[1~128]个字符。
	Name *string `json:"name,omitempty"`

	// **参数解释：** 规则集中文描述。 **取值范围：** 不涉及。
	Comments *string `json:"comments,omitempty"`

	// **参数解释：** 规则集类型。 **取值范围：** - system：系统规则集。 - custom：自定义规则集。
	Type *string `json:"type,omitempty"`

	// **参数解释：** 规则集创建者。 **取值范围：** 用户名。
	Creator *string `json:"creator,omitempty"`

	// **参数解释：** 规则集创建时间。 **取值范围：** 0时区标准时间。
	CreateTime *string `json:"createTime,omitempty"`

	// **参数解释：** 规则集最近一次更新的操作人。 **取值范围：** 用户名。
	LastReviser *string `json:"lastReviser,omitempty"`

	// **参数解释：** 规则集最近一次更新的时间。 **取值范围：** 0时区标准时间。
	LastUpdateTime *string `json:"lastUpdateTime,omitempty"`

	// **参数解释：** 规则集语言。 **取值范围：** - ARKTS - C# - CANGJIE - C++ - CSS - GO - HTML - JAVA - JAVASCRIPT - KOTLIN - LUA - PHP - PYTHON - RUST - SCALA - SHELL - SQL - TERRAFORM - TYPESCRIPT - AZURERESOURCEMANAGER - DOCKER - KUBERNETES
	Language *string `json:"language,omitempty"`

	// **参数解释：** 父规则集ID。 **取值范围：** 32位UUID组成的字符串。多个ID之间用逗号分隔。
	ParentId *string `json:"parentId,omitempty"`

	// **参数解释：** 规则集中已启用的规则数量。 **取值范围：** ≥0
	NumCriterion *int32 `json:"numCriterion,omitempty"`

	// **参数解释：** 当前用户是否具有修改该规则集的权限。 **取值范围：** - true：有编辑的权限。 - false：没有编辑的权限。
	CanEdit *bool `json:"canEdit,omitempty"`

	// **参数解释：** 当前用户是否具有删除该规则集的权限。 **取值范围：** - true：有删除的权限。 - false：没有删除的权限。
	CanDelete *bool `json:"canDelete,omitempty"`

	// **参数解释：** 规则集管理员。 **取值范围：** 用户ID。多个ID用逗号分隔。
	Managers *string `json:"managers,omitempty"`

	// **参数解释：** 父规则集的名称。 **取值范围：** 多个规则集名称之间用逗号分隔。
	ParentSetNames *string `json:"parentSetNames,omitempty"`

	// **参数解释：** 租户ID，租户ID的值来源于IAM的租户ID。 **取值范围：** 32位UUID。
	DomainId *string `json:"domainId,omitempty"`

	// **参数解释：** 规则集创建时候的继承类型。 **取值范围：** - strong：规则集创建的时候来源于继承。 - null：规则集创建的时候来源于非继承。
	InheritMode *string `json:"inheritMode,omitempty"`

	// **参数解释：** 创建者用户ID。 **取值范围：** 32位UUID。
	CreatorId *string `json:"creatorId,omitempty"`

	// **参数解释：** 是否为系统默认规则集。 **取值范围：** - 0：否。 - 1：是。
	IsSystemDefault *string `json:"isSystemDefault,omitempty"`

	// **参数解释：** 是否为项目默认规则集。 **取值范围：** - 0：否。 - 1：是。
	IsProjectDefault *string `json:"isProjectDefault,omitempty"`

	// **参数解释：** 是否为系统规则集。 **取值范围：** - 0：否。 - 1：是。
	IsSystemSet *string `json:"isSystemSet,omitempty"`

	// **参数解释：** 规则集中已启用的能力规则unique_id。 **取值范围：** unique_id为32位UUID。多个unique_id之间用逗号分隔。
	RuleIds *string `json:"ruleIds,omitempty"`

	// **参数解释：** 规则集所在局点。 **取值范围：** 不涉及。
	Region *string `json:"region,omitempty"`

	// **参数解释：** 是否具有设置默认规则集的权限。 **取值范围：** - true：有权限。 - false：没有权限。
	CanSetDefault *bool `json:"canSetDefault,omitempty"`

	// **参数解释：** 规则集是否被使用。 **取值范围：** - true：规则集被使用。 - false：规则集未被使用。
	IsUsed *bool `json:"isUsed,omitempty"`

	// **参数解释：** 子规则集。 **取值范围：** 不涉及。
	ChildSets *[]string `json:"childSets,omitempty"`

	// **参数解释：** 是否为可信规则集。该字段已弃用。 **取值范围：** - true：是。 - false：否。
	IsReliable *bool `json:"isReliable,omitempty"`

	// **参数解释：** 是否为推荐规则集。该字段已弃用。 **取值范围：** - 1：是。 - 0：否。
	IsRecommand *string `json:"isRecommand,omitempty"`

	// **参数解释：** 是否为默认规则集。该字段已弃用。 **取值范围：** - true：是。 - false：否。
	IsDefault *bool `json:"isDefault,omitempty"`

	// **参数解释：** 是否为默认规则集。该字段已弃用。 **取值范围：** - \"true\"：是。 - \"false\"：否。
	DefaultSet *string `json:"defaultSet,omitempty"`

	// **参数解释：** 该字段已弃用。 **取值范围：** 不涉及。
	DeptId *string `json:"deptId,omitempty"`

	// **参数解释：** 该字段已弃用。 **取值范围：** 不涉及。
	DeptName *string `json:"deptName,omitempty"`

	// **参数解释：** 该字段已弃用。 **取值范围：** 不涉及。
	ReliableTag *string `json:"reliableTag,omitempty"`

	// **参数解释：** 该字段已弃用。 **取值范围：** 不涉及。
	RecommandType *string `json:"recommandType,omitempty"`
}

func (o CriterionSetListV2Body) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CriterionSetListV2Body struct{}"
	}

	return strings.Join([]string{"CriterionSetListV2Body", string(data)}, " ")
}
