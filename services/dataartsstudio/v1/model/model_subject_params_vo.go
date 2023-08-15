package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SubjectParamsVo struct {

	// 编码。更新时必填，创建时可以为空
	Id *int64 `json:"id,omitempty"`

	// 中文名称
	NameCh string `json:"name_ch"`

	// 英文名称
	NameEn string `json:"name_en"`

	// 描述信息, 业务对象必填
	Description *string `json:"description,omitempty"`

	// 别名
	Alias *string `json:"alias,omitempty"`

	// 数据owner部门
	DataOwner *string `json:"data_owner,omitempty"`

	// 数据owner人员
	DataOwnerList string `json:"data_owner_list"`

	// 层级
	Level int32 `json:"level"`

	// 上层主题id，首层则为空
	ParentId *int64 `json:"parent_id,omitempty"`

	// 属性自定义项
	SelfDefinedFields *[]SelfDefinedFieldVo `json:"self_defined_fields,omitempty"`
}

func (o SubjectParamsVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubjectParamsVo struct{}"
	}

	return strings.Join([]string{"SubjectParamsVo", string(data)}, " ")
}
