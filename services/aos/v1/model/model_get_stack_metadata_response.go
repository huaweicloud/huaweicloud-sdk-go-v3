package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// GetStackMetadataResponse Response Object
type GetStackMetadataResponse struct {

	// 资源栈（stack）的唯一ID。  此ID由资源编排服务在生成资源栈的时候生成，为UUID。  由于资源栈名仅仅在同一时间下唯一，即用户允许先生成一个叫HelloWorld的资源栈，删除，再重新创建一个同名资源栈。  对于团队并行开发，用户可能希望确保，当前我操作的资源栈就是我认为的那个，而不是其他队友删除后创建的同名资源栈。因此，使用ID就可以做到强匹配。  资源编排服务保证每次创建的资源栈所对应的ID都不相同，更新不会影响ID。如果给予的stack_id和当前资源栈的ID不一致，则返回400
	StackId *string `json:"stack_id,omitempty"`

	// 资源栈的名称。此名字在domain_id+区域+project_id下应唯一，可以使用中文、大小写英文、数字、下划线、中划线。首字符需为中文或者英文，区分大小写。
	StackName string `json:"stack_name"`

	// 资源栈的描述。可用于客户识别自己的资源栈。
	Description *string `json:"description,omitempty"`

	// HCL参数结构。HCL模板支持参数传入，即，同一个模板可以给予不同的参数而达到不同的效果。  * var_structure可以允许客户提交最简单的字符串类型的参数  * 资源编排服务支持vars_structure，vars_body和vars_uri，如果以上三种方式中声名了同一个变量，将报错400  * vars_structure中的值只支持简单的字符串类型，如果需要使用其他类型，需要用户自己在HCL引用时转换， 或者用户可以使用vars_uri、vars_body，vars_uri和vars_body中支持HCL支持的各种类型以及复杂结构  * 如果vars_structure过大，可以使用vars_uri  * 注意：vars_structure中默认不应该含有任何敏感信息，资源编排服务会直接明文使用、log、展示、存储对应的vars。如为敏感信息，建议设置encryption字段开启加密
	VarsStructure *[]VarsStructure `json:"vars_structure,omitempty"`

	// HCL参数文件的内容。HCL模板支持参数传入，即，同一个模板可以给予不同的参数而达到不同的效果。  * vars_body使用HCL的tfvars格式，用户可以将“.tfvars”中的内容提交到vars_body中  * 资源编排服务支持vars_body和vars_uri，如果以上两种方式中声名了同一个变量，将报错400  * 如果vars_body过大，可以使用vars_uri  * 资源栈集不支持敏感数据加密，资源编排服务会直接明文使用、log、展示、存储对应的vars_body。
	VarsBody *string `json:"vars_body,omitempty"`

	// 删除保护的标识位，如果不传默认为false，即默认不开启资源栈删除保护（删除保护开启后资源栈不允许被删除）  *在UpdateStack API中，如果该参数未在RequestBody中给予，则不会对资源栈的删除保护属性进行更新*
	EnableDeletionProtection *bool `json:"enable_deletion_protection,omitempty"`

	// 自动回滚的标识位，如果不传默认为false，即默认不开启资源栈自动回滚（自动回滚开启后，如果部署失败，则会自动回滚，并返回上一个稳定状态）  *在UpdateStack API中，如果该参数未在RequestBody中给予，则不会对资源栈的自动回滚属性进行更新* *该属性与使用模板导入资源功能互斥，如果堆栈的自动回滚设置为true，则不允许部署包含导入资源的模板*
	EnableAutoRollback *bool `json:"enable_auto_rollback,omitempty"`

	// 资源栈的状态    * `CREATION_COMPLETE` - 生成空资源栈完成，并没有任何部署    * `DEPLOYMENT_IN_PROGRESS` - 正在部署，请等待    * `DEPLOYMENT_FAILED` - 部署失败。请从status_message获取错误信息汇总，或者调用ListStackEvents获得事件详情    * `DEPLOYMENT_COMPLETE` - 部署完成    * `ROLLBACK_IN_PROGRESS` - 部署失败，正在回滚，请等待    * `ROLLBACK_FAILED` - 回滚失败。请从status_message获取错误信息汇总，或者调用ListStackEvents获得事件详情    * `ROLLBACK_COMPLETE` - 回滚完成    * `DELETION_IN_PROGRESS` - 正在删除，请等待    * `DELETION_FAILED` - 删除失败。请从status_message获取错误信息汇总，或者调用ListStackEvents获得事件详情
	Status *GetStackMetadataResponseStatus `json:"status,omitempty"`

	// 委托授权的信息。  RFS仅在创建资源栈（触发部署）、创建执行计划、部署资源栈、删除资源栈等涉及资源操作的请求中使用委托，且该委托仅作用于与之绑定的Provider对资源的操作中。如果委托中提供的权限不足，有可能导致相关资源操作失败。  [[创建委托及授权方式](https://support.huaweicloud.com/usermanual-iam/iam_06_0002.html)](tag:hws) [[创建委托及授权方式](https://support.huaweicloud.com/intl/zh-cn/usermanual-iam/iam_06_0002.html)](tag:hws_hk) [[创建委托及授权方式](https://support.huaweicloud.com/eu/usermanual-iam/iam_06_0002.html)](tag:hws_eu)
	Agencies *[]Agency `json:"agencies,omitempty"`

	// 当资源栈的状态为任意失败状态（即以 `FAILED` 结尾时），将会展示简要的错误信息总结以供debug
	StatusMessage *string `json:"status_message,omitempty"`

	// vars_uri对应的文件内容
	VarsUriContent *string `json:"vars_uri_content,omitempty"`

	// 资源栈的生成时间 格式遵循RFC3339，即yyyy-mm-ddTHH:MM:SSZ，如1970-01-01T00:00:00Z
	CreateTime *string `json:"create_time,omitempty"`

	// 资源栈的更新时间（更新场景包括元数据更新场景和部署场景） 格式遵循RFC3339，即yyyy-mm-ddTHH:MM:SSZ，如1970-01-01T00:00:00Z
	UpdateTime     *string `json:"update_time,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o GetStackMetadataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetStackMetadataResponse struct{}"
	}

	return strings.Join([]string{"GetStackMetadataResponse", string(data)}, " ")
}

type GetStackMetadataResponseStatus struct {
	value string
}

type GetStackMetadataResponseStatusEnum struct {
	CREATION_COMPLETE      GetStackMetadataResponseStatus
	DEPLOYMENT_IN_PROGRESS GetStackMetadataResponseStatus
	DEPLOYMENT_FAILED      GetStackMetadataResponseStatus
	DEPLOYMENT_COMPLETE    GetStackMetadataResponseStatus
	ROLLBACK_IN_PROGRESS   GetStackMetadataResponseStatus
	ROLLBACK_FAILED        GetStackMetadataResponseStatus
	ROLLBACK_COMPLETE      GetStackMetadataResponseStatus
	DELETION_IN_PROGRESS   GetStackMetadataResponseStatus
	DELETION_FAILED        GetStackMetadataResponseStatus
}

func GetGetStackMetadataResponseStatusEnum() GetStackMetadataResponseStatusEnum {
	return GetStackMetadataResponseStatusEnum{
		CREATION_COMPLETE: GetStackMetadataResponseStatus{
			value: "CREATION_COMPLETE",
		},
		DEPLOYMENT_IN_PROGRESS: GetStackMetadataResponseStatus{
			value: "DEPLOYMENT_IN_PROGRESS",
		},
		DEPLOYMENT_FAILED: GetStackMetadataResponseStatus{
			value: "DEPLOYMENT_FAILED",
		},
		DEPLOYMENT_COMPLETE: GetStackMetadataResponseStatus{
			value: "DEPLOYMENT_COMPLETE",
		},
		ROLLBACK_IN_PROGRESS: GetStackMetadataResponseStatus{
			value: "ROLLBACK_IN_PROGRESS",
		},
		ROLLBACK_FAILED: GetStackMetadataResponseStatus{
			value: "ROLLBACK_FAILED",
		},
		ROLLBACK_COMPLETE: GetStackMetadataResponseStatus{
			value: "ROLLBACK_COMPLETE",
		},
		DELETION_IN_PROGRESS: GetStackMetadataResponseStatus{
			value: "DELETION_IN_PROGRESS",
		},
		DELETION_FAILED: GetStackMetadataResponseStatus{
			value: "DELETION_FAILED",
		},
	}
}

func (c GetStackMetadataResponseStatus) Value() string {
	return c.value
}

func (c GetStackMetadataResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GetStackMetadataResponseStatus) UnmarshalJSON(b []byte) error {
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
