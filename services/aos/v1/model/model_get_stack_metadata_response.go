package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type GetStackMetadataResponse struct {

	// 栈的唯一Id
	StackId *string `json:"stack_id,omitempty"`

	// 栈的名字
	StackName *string `json:"stack_name,omitempty"`

	// 栈的描述，此描述为用户在创建资源栈时指定
	Description *string `json:"description,omitempty"`

	// 参数列表
	VarsStructure *[]VarsStructure `json:"vars_structure,omitempty"`

	// vars文件中的内容
	VarsUriContent *string `json:"vars_uri_content,omitempty"`

	// terraform支持参数，即，同一个模板可以给予不同的参数而达到不同的效果。vars_body用于接收用户直接提交的tfvars文件内容
	VarsBody *string `json:"vars_body,omitempty"`

	// 栈的生成时间，格式遵循RFC3339，即yyyy-mm-ddTHH:MM:SSZ，如1970-01-01T00:00:00Z
	CreateTime *string `json:"create_time,omitempty"`

	// 由于栈可以被更新，此处为上次更新时间，格式遵循RFC3339，即yyyy-mm-ddTHH:MM:SSZ，如1970-01-01T00:00:00Z
	UpdateTime *string `json:"update_time,omitempty"`

	// 资源栈删除保护的目标状态
	EnableDeletionProtection *bool `json:"enable_deletion_protection,omitempty"`

	// 资源栈是否开启自动回滚的标识位
	EnableAutoRollback *bool `json:"enable_auto_rollback,omitempty"`

	// 资源栈的执行状态     * `DEPLOYMENT_IN_PROGRESS` - 正在部署     * `DEPLOYMENT_FAILED` - 部署失败。请于StatusMessage见更多详情     * `DEPLOYMENT_COMPLETE ` - 部署结束     * `ROLLBACK_IN_PROGRESS` - 正在回滚     * `ROLLBACK_FAILED` - 回滚失败。请于StatusMessage见更多详情     * `ROLLBACK_COMPLETE` - 回滚完成     * `DELETION_IN_PROGRESS` - 正在删除     * `DELETION_FAILED` - 删除失败     * `CREATION_COMPLETE` - 生成完成，并没有任何部署
	Status *GetStackMetadataResponseStatus `json:"status,omitempty"`

	// 展示更多细节的信息
	StatusMessage *string `json:"status_message,omitempty"`

	// 委托授权的信息
	Agencies       *[]Agency `json:"agencies,omitempty"`
	HttpStatusCode int       `json:"-"`
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
	DEPLOYMENT_IN_PROGRESS GetStackMetadataResponseStatus
	DEPLOYMENT_FAILED      GetStackMetadataResponseStatus
	DEPLOYMENT_COMPLETE    GetStackMetadataResponseStatus
	ROLLBACK_IN_PROGRESS   GetStackMetadataResponseStatus
	ROLLBACK_FAILED        GetStackMetadataResponseStatus
	ROLLBACK_COMPLETE      GetStackMetadataResponseStatus
	DELETION_IN_PROGRESS   GetStackMetadataResponseStatus
	DELETION_FAILED        GetStackMetadataResponseStatus
	CREATION_COMPLETE      GetStackMetadataResponseStatus
}

func GetGetStackMetadataResponseStatusEnum() GetStackMetadataResponseStatusEnum {
	return GetStackMetadataResponseStatusEnum{
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
		CREATION_COMPLETE: GetStackMetadataResponseStatus{
			value: "CREATION_COMPLETE",
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
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
