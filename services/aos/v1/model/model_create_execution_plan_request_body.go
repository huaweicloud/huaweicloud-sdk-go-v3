package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// create-execution-plan request parameters
type CreateExecutionPlanRequestBody struct {

	// 用户希望生成执行计划的栈（stack）的Id。此Id由资源编排服务在生成栈的时候生成，为UUID。
	StackId *string `json:"stack_id,omitempty"`

	// HCL模板，描述了资源的目标状态。资源编排服务将比较此模板与当前远程资源的状态之间的区别 template_body 和 template_uri 有且仅有一个存在
	TemplateBody *string `json:"template_body,omitempty"`

	// HCL模板的OBS地址，描述了资源的目标状态。资源编排服务将比较此模板与当前远程资源的状态之间的区别。目前接受纯tf文件或zip压缩包。 纯tf文件需要以“.tf”结尾，并遵守tf模板格式。压缩包目前只支持zip格式，文件需要以\".zip\"结尾，压缩包解压后应该只包含文件，且文件均以“.tf”结尾，不支持nested结构 template_body 和 template_uri 有且仅有一个存在
	TemplateUri *string `json:"template_uri,omitempty"`

	// 执行计划的名字，在domain_id+region+project_id+stack_id下应唯一。
	ExecutionPlanName string `json:"execution_plan_name"`

	// 执行计划的描述。可用于客户识别自己的执行计划
	Description *string `json:"description,omitempty"`

	// terraform支持参数，即，同一个模板可以给予不同的参数而达到不同的效果。var是一系列terraform所需要的参数。 注：资源编排服务支持vars、vars_body和vars_uri，如果vars、vars_body和vars_uri中声名了同一个变量，将报错400。 注：vars中的值只支持简单的字符串类型，如果其他类型，需要用户自己在HCL引用时转换，或者用户可以使用vars_body或vars_uri， vars_body和vars_uri中支持HCL支持的各种类型以及复杂结构。
	VarsStructure *[]VarsStructure `json:"vars_structure,omitempty"`

	// terraform支持参数，即，同一个模板可以给予不同的参数而达到不同的效果。vars_body用于接收用户直接提交的tfvars文件内容
	VarsBody *string `json:"vars_body,omitempty"`

	// 参数文件的OBS地址，如果客户偏向使用文件维护参数，可以将参数上传OBS，并将OBS地址提交。 注：如果用户同时使用了vars_body、vars_uri和vars，且他们的内容中定义了同一个参数，资源编排服务将报错并返回400。 vars_uri和vars_body中的vars按照HCL的语义，可以支持各种类型、复杂结构等
	VarsUri *string `json:"vars_uri,omitempty"`
}

func (o CreateExecutionPlanRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateExecutionPlanRequestBody struct{}"
	}

	return strings.Join([]string{"CreateExecutionPlanRequestBody", string(data)}, " ")
}
