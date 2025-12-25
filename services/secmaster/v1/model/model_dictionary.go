package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Dictionary struct {

	// uuid
	Id *string `json:"id,omitempty"`

	// 版本号
	Version *string `json:"version,omitempty"`

	// 字典id
	DictId *string `json:"dict_id,omitempty"`

	// 字典key
	DictKey *string `json:"dict_key,omitempty"`

	// 字典code
	DictCode *string `json:"dict_code,omitempty"`

	// 字典值
	DictVal *string `json:"dict_val,omitempty"`

	// 字典关联的父key
	DictPkey *string `json:"dict_pkey,omitempty"`

	// 字典关联的父code
	DictPcode *string `json:"dict_pcode,omitempty"`

	// 创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 更新时间
	UpdateTime *string `json:"update_time,omitempty"`

	// 发布时间
	PublishTime *string `json:"publish_time,omitempty"`

	// 字典所属领域
	Scope *string `json:"scope,omitempty"`

	// 字典描述信息
	Description *string `json:"description,omitempty"`

	// 额外字段
	ExtensionField *interface{} `json:"extension_field,omitempty"`

	// **参数解释：** 项目ID，用于明确项目归属，配置后可通过该ID查询项目下资产，可以通过调用API获取，也可以从控制台获取。[获取项目ID](secmaster_03_0014.xml) **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	ProjectId *string `json:"project_id,omitempty"`

	// 用户当前语言环境
	Language *string `json:"language,omitempty"`
}

func (o Dictionary) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Dictionary struct{}"
	}

	return strings.Join([]string{"Dictionary", string(data)}, " ")
}
