package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateStackRequestBody struct {

	// 用户希望生成的资源栈的名字。此名字在domain_id+区域+project_id下应唯一，可以使用中文、大小写英文、数字、下划线、中划线。首字符需为中文或者英文，区分大小写。
	StackName string `json:"stack_name"`

	// 委托授权的信息。
	Agencies *[]Agency `json:"agencies,omitempty"`

	// 资源栈的描述。可用于客户识别自己的资源栈。
	Description *string `json:"description,omitempty"`

	// 删除保护的标识位，如果不传默认为false，即默认不开启资源栈删除保护（删除保护开启后资源栈不允许被删除）
	EnableDeletionProtection *bool `json:"enable_deletion_protection,omitempty"`

	// 自动回滚的标识位，如果不传默认为false，即默认不开启资源栈自动回滚（自动回滚开启后，如果部署失败，则会自动回滚，并返回上一个稳定状态）
	EnableAutoRollback *bool `json:"enable_auto_rollback,omitempty"`

	// HCL模板，描述了资源的目标状态。资源编排服务将比较此模板与当前远程资源的状态之间的区别。  template_body和template_uri 必须有且只有一个存在
	TemplateBody *string `json:"template_body,omitempty"`

	// HCL模板的OBS地址，该模板描述了资源的目标状态  OBS地址必须为同region的OBS地址，暂不支持跨region访问  对应的文件应该是纯tf文件或zip压缩包  纯tf文件需要以`.tf`或者`.tfjson`结尾，并遵守hcl语法  压缩包目前只支持zip格式，文件需要以\".zip\"结尾。解压后的文件不得包含\".tfvars\"文件  template_body和template_uri 必须有且只有一个存在
	TemplateUri *string `json:"template_uri,omitempty"`

	// HCL支持参数，即，同一个模板可以给予不同的参数而达到不同的效果  * vars_body使用HCL的tfvars格式，用户可以将“.tfvars”中的内容提交到vars_body中。具体tfvars格式见：https://www.terraform.io/language/values/variables#variable-definitions-tfvars-files  * 资源编排服务支持vars_structure，vars_body和vars_uri，如果他们中声名了同一个变量，将报错400  * 如果vars_body过大，可以使用vars_uri  * 如果vars中都是简单的字符串格式，可以使用var_structure  * 注意：vars中不应该传递任何敏感信息，资源编排服务会直接明文使用、log、展示、存储对应的vars
	VarsBody *string `json:"vars_body,omitempty"`

	// HCL支持参数，即，同一个模板可以给予不同的参数而达到不同的效果。  * var_structure可以允许客户提交最简单的字符串类型的参数  * 资源编排服务支持vars_structure，vars_body和vars_uri，如果他们中声名了同一个变量，将报错400  * vars_structure中的值只支持简单的字符串类型，如果需要使用其他类型，需要用户自己在HCL引用时转换， 或者用户可以使用vars_uri、vars_body，vars_uri和vars_body中支持HCL支持的各种类型以及复杂结构  * 如果vars_structure过大，可以使用vars_uri  * 注意：vars中不应该传递任何敏感信息，资源编排服务会直接明文使用、log、展示、存储对应的vars
	VarsStructure *[]VarsStructure `json:"vars_structure,omitempty"`

	// HCL支持参数，即，同一个模板可以给予不同的参数而达到不同的效果  * vars_body使用HCL的tfvars格式，用户可以将“.tfvars”中的内容提交到vars_body中。具体tfvars格式见：https://www.terraform.io/language/values/variables#variable-definitions-tfvars-files  * 资源编排服务支持vars_structure，vars_body和vars_uri，如果他们中声名了同一个变量，将报错400  * 如果vars_body过大，可以使用vars_uri  * 如果vars中都是简单的字符串格式，可以使用var_structure  * 注意：vars中不应该传递任何敏感信息，资源编排服务会直接明文使用、log、展示、存储对应的vars
	VarsUri *string `json:"vars_uri,omitempty"`
}

func (o CreateStackRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateStackRequestBody struct{}"
	}

	return strings.Join([]string{"CreateStackRequestBody", string(data)}, " ")
}
