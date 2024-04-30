package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// GetExecutionPlanMetadataResponse Response Object
type GetExecutionPlanMetadataResponse struct {

	// 资源栈（stack）的唯一ID。  此ID由资源编排服务在生成资源栈的时候生成，为UUID。  由于资源栈名仅仅在同一时间下唯一，即用户允许先生成一个叫HelloWorld的资源栈，删除，再重新创建一个同名资源栈。  对于团队并行开发，用户可能希望确保，当前我操作的资源栈就是我认为的那个，而不是其他队友删除后创建的同名资源栈。因此，使用ID就可以做到强匹配。  资源编排服务保证每次创建的资源栈所对应的ID都不相同，更新不会影响ID。如果给予的stack_id和当前资源栈的ID不一致，则返回400
	StackId *string `json:"stack_id,omitempty"`

	// 资源栈的名称。此名字在domain_id+区域+project_id下应唯一，可以使用中文、大小写英文、数字、下划线、中划线。首字符需为中文或者英文，区分大小写。
	StackName string `json:"stack_name"`

	// 执行计划（execution_plan）的唯一Id。  此Id由资源编排服务在生成执行计划的时候生成，为UUID。  由于执行计划名仅仅在同一时间下唯一，即用户允许先生成一个叫HelloWorld的执行计划，删除，再重新创建一个同名执行计划。  对于团队并行开发，用户可能希望确保，当前我操作的执行计划就是我认为的那个，而不是其他队友删除后创建的同名执行计划。因此，使用ID就可以做到强匹配。  资源编排服务保证每次创建的执行计划所对应的ID都不相同，更新不会影响ID。如果给予的execution_plan_id和当前执行计划的ID不一致，则返回400  **注意：** * 创建执行计划后，资源编排服务持久化请求并立即返回，客户端不等待请求最终处理完成，用户无法实时感知请求处理结果 * 资源编排服务最终会将异步部署请求排队，在服务端空闲的情况下逐个处理。用户最大等待时长为1小时
	ExecutionPlanId *string `json:"execution_plan_id,omitempty"`

	// 执行计划的名称。此名字在domain_id+区域+project_id+stack_id下应唯一，可以使用中文、大小写英文、数字、下划线、中划线。首字符需为中文或者英文，区分大小写。
	ExecutionPlanName string `json:"execution_plan_name"`

	// 执行计划的描述。可用于客户识别自己的执行计划。
	Description *string `json:"description,omitempty"`

	// HCL参数结构。HCL模板支持参数传入，即，同一个模板可以给予不同的参数而达到不同的效果。  * var_structure可以允许客户提交最简单的字符串类型的参数  * 资源编排服务支持vars_structure，vars_body和vars_uri，如果以上三种方式中声名了同一个变量，将报错400  * vars_structure中的值只支持简单的字符串类型，如果需要使用其他类型，需要用户自己在HCL引用时转换， 或者用户可以使用vars_uri、vars_body，vars_uri和vars_body中支持HCL支持的各种类型以及复杂结构  * 如果vars_structure过大，可以使用vars_uri  * 注意：vars_structure中默认不应该含有任何敏感信息，资源编排服务会直接明文使用、log、展示、存储对应的vars。如为敏感信息，建议设置encryption字段开启加密
	VarsStructure *[]VarsStructure `json:"vars_structure,omitempty"`

	// vars_uri对应的文件内容
	VarsUriContent *string `json:"vars_uri_content,omitempty"`

	// HCL参数文件的内容。HCL模板支持参数传入，即，同一个模板可以给予不同的参数而达到不同的效果。  * vars_body使用HCL的tfvars格式，用户可以将“.tfvars”中的内容提交到vars_body中  * 资源编排服务支持vars_body和vars_uri，如果以上两种方式中声名了同一个变量，将报错400  * 如果vars_body过大，可以使用vars_uri  * 资源栈集不支持敏感数据加密，资源编排服务会直接明文使用、log、展示、存储对应的vars_body。
	VarsBody *string `json:"vars_body,omitempty"`

	// 执行计划的状态    * `CREATION_IN_PROGRESS` - 正在创建，请等待    * `CREATION_FAILED` - 创建失败，请从status_message获取错误信息汇总    * `AVAILABLE` - 创建完成，可以调用ApplyExecutionPlan API进行执行    * `APPLY_IN_PROGRESS` - 执行中，可通过GetStackMetadata查询资源栈状态，通过ListStackEvents获取执行过程中产生的资源栈事件    * `APPLIED` - 已执行
	Status *GetExecutionPlanMetadataResponseStatus `json:"status,omitempty"`

	// 当执行计划的状态为创建失败状态（即为 `CREATION_FAILED` 时），将会展示简要的错误信息总结以供debug
	StatusMessage *string `json:"status_message,omitempty"`

	// 执行计划的生成时间 格式遵循RFC3339，即yyyy-mm-ddTHH:MM:SSZ，如1970-01-01T00:00:00Z
	CreateTime *string `json:"create_time,omitempty"`

	// 执行计划的执行时间 格式遵循RFC3339，即yyyy-mm-ddTHH:MM:SSZ，如1970-01-01T00:00:00Z
	ApplyTime *string `json:"apply_time,omitempty"`

	Summary        *ExecutionPlanSummary `json:"summary,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o GetExecutionPlanMetadataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetExecutionPlanMetadataResponse struct{}"
	}

	return strings.Join([]string{"GetExecutionPlanMetadataResponse", string(data)}, " ")
}

type GetExecutionPlanMetadataResponseStatus struct {
	value string
}

type GetExecutionPlanMetadataResponseStatusEnum struct {
	CREATION_IN_PROGRESS GetExecutionPlanMetadataResponseStatus
	CREATION_FAILED      GetExecutionPlanMetadataResponseStatus
	AVAILABLE            GetExecutionPlanMetadataResponseStatus
	APPLY_IN_PROGRESS    GetExecutionPlanMetadataResponseStatus
	APPLIED              GetExecutionPlanMetadataResponseStatus
}

func GetGetExecutionPlanMetadataResponseStatusEnum() GetExecutionPlanMetadataResponseStatusEnum {
	return GetExecutionPlanMetadataResponseStatusEnum{
		CREATION_IN_PROGRESS: GetExecutionPlanMetadataResponseStatus{
			value: "CREATION_IN_PROGRESS",
		},
		CREATION_FAILED: GetExecutionPlanMetadataResponseStatus{
			value: "CREATION_FAILED",
		},
		AVAILABLE: GetExecutionPlanMetadataResponseStatus{
			value: "AVAILABLE",
		},
		APPLY_IN_PROGRESS: GetExecutionPlanMetadataResponseStatus{
			value: "APPLY_IN_PROGRESS",
		},
		APPLIED: GetExecutionPlanMetadataResponseStatus{
			value: "APPLIED",
		},
	}
}

func (c GetExecutionPlanMetadataResponseStatus) Value() string {
	return c.value
}

func (c GetExecutionPlanMetadataResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GetExecutionPlanMetadataResponseStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
