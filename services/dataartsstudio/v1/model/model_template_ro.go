package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TemplateRo struct {

	// template id
	Id *string `json:"id,omitempty"`

	// template name
	Name *string `json:"name,omitempty"`

	// 目录
	DirectoryId *int64 `json:"directory_id,omitempty"`

	// 维度ID, 1:完整性,2:唯一性,3:及时性,4:有效性,5:准确性,6:一致性
	DimensionId *int32 `json:"dimension_id,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 模板中的sql内容
	SqlInfo *string `json:"sql_info,omitempty"`

	// 结果说明
	ResultDescription *string `json:"result_description,omitempty"`

	// 是否是发布操作， true：发布新版本  false：普通的保存操作
	Publish *bool `json:"publish,omitempty"`

	// 修改前的模板名
	OriginName *string `json:"origin_name,omitempty"`

	// 异常表模板
	AbnormalTableTemplate *string `json:"abnormal_table_template,omitempty"`

	// 用户自定义版本名
	UserDefineVersionName *string `json:"user_define_version_name,omitempty"`

	// 获取模板信息时候的版本号
	VersionNum *int64 `json:"version_num,omitempty"`

	// 规则模板状态,0表示下线1表示已发布
	Status *int32 `json:"status,omitempty"`
}

func (o TemplateRo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TemplateRo struct{}"
	}

	return strings.Join([]string{"TemplateRo", string(data)}, " ")
}
