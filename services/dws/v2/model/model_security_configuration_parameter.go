package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SecurityConfigurationParameter 安全设置参数响应体。
type SecurityConfigurationParameter struct {

	// **参数解释**： 参数名。 **取值范围**： audit_dml_state：审计dml操作开关； audit_system_object：审计DDL操作、其它操作； audit_adm：安全管理员用户名； audit_exec_status：审计执行结果； audit_operation_checked：审计DML操作、审计其它操作的具体勾选项； enableSeparationOfDuty：三权分立开关； audit_user_violation：越权访问操作； ssl：ssl开关； require_ssl：是否校验ssl； audit_function_exec：审计存储过程执行操作； audit_copy_exec：对COPY操作进行记录； audit_resource_policy：日志保留策略； audit_file_remain_time：时间策略下的最少保留天数，已废弃； audit_dml_state_select：审计SELECT操作； security_adm：安全管理员； audit_dump_switch：日志转储开关； kernel_audit_dump_switch：内核日志转储开关； audit_system_object_detail：审计DDL操作、其它操作的具体勾选项；
	Name *string `json:"name,omitempty"`

	// **参数解释**： 参数值。 **取值范围**： 不涉及。
	Value *string `json:"value,omitempty"`
}

func (o SecurityConfigurationParameter) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SecurityConfigurationParameter struct{}"
	}

	return strings.Join([]string{"SecurityConfigurationParameter", string(data)}, " ")
}
