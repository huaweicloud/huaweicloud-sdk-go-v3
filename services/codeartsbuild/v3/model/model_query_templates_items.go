package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QueryTemplatesItems struct {

	// 是否收藏模板
	Favourite *bool `json:"favourite,omitempty"`

	// 昵称
	NickName *string `json:"nick_name,omitempty"`

	// 数据库中id
	Id *string `json:"id,omitempty"`

	// uuid
	Uuid *string `json:"uuid,omitempty"`

	Template *QueryTemplateVo `json:"template,omitempty"`

	// 模板类别
	Type *string `json:"type,omitempty"`

	// 模板是否公开
	Public *bool `json:"public,omitempty"`

	// 模板命名
	Name *string `json:"name,omitempty"`

	// 创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// domainId
	DomainId *string `json:"domain_id,omitempty"`

	// 权重
	Weight *float64 `json:"weight,omitempty"`

	// 用户id
	UserId *string `json:"user_id,omitempty"`

	// 用户名
	UserName *string `json:"user_name,omitempty"`

	// domain名字
	DomainName *string `json:"domain_name,omitempty"`

	// 模板范围，自定义模板默认为custom
	Scope *string `json:"scope,omitempty"`

	// 模板说明
	Description *string `json:"description,omitempty"`

	// 构建工具类型，yaml构建还是action构建
	ToolType *string `json:"tool_type,omitempty"`

	// intl说明
	IntlDescription *interface{} `json:"intl_description,omitempty"`

	// 构建执行参数列表
	Parameters *[]CreateBuildJobParameter `json:"parameters,omitempty"`

	// i18n
	I18n *interface{} `json:"i18n,omitempty"`
}

func (o QueryTemplatesItems) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryTemplatesItems struct{}"
	}

	return strings.Join([]string{"QueryTemplatesItems", string(data)}, " ")
}
