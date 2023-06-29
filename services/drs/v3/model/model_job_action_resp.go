package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type JobActionResp struct {

	// 任务可操作命令集合。 取值： CREATE：创建任务 CHOOSE_OBJECT：选择对象 PRE_CHECK：预检查 CHANGE_MODE：修改任务模式 FREE_RESOURCE：释放资源 MODIFY_DB_CONFIG：修改数据库配置 RESET_DB_PWD：重置数据库密码（源库、目标库） MODIFY_CONFIGURATION：修改任务配置 PAUSE：暂停任务 START：启动任务 CHANGE：修改任务 RETRY：重试任务 RESET：重置任务 DELETE：删除任务 QUERY_PRE_CHECK：预检查 SWITCH_OVER：容灾倒换 START_INCR：CASSANDRA启动增量任务 MODIFY_TASK_NUMBER：CASSANDRA修改线程数配置 CONTINUE_JOB：oracle-GaussDB分布式:启动失败或者停止的任务 STOP_JOB：oracle-GaussDB分布式:停止任务 CONTINUE_CAPTURE：oracle-GaussDB分布式:启动抓取 STOP_CAPTURE：oracle-GaussDB分布式:停止抓取 CONTINUE_APPLY：oracle-GaussDB分布式:启动回放 STOP_APPLY：oracle-GaussDB分布式:停止回放 PAY_ORDER：包年包月支付订单 UNSUBSCRIBE：包年包月退订 TO_PERIOD：转包周期 TO_RENEW：包周期续费 ORDER_INFO：订单详情 CHANGE_FLAVOR：规格变更 CLONE：克隆任务
	AvailableActions *[]JobActionRespAvailableActions `json:"available_actions,omitempty"`

	// 任务不可操作命令集合。 取值： CREATE：创建任务 CHOOSE_OBJECT：选择对象 PRE_CHECK：预检查 CHANGE_MODE：修改任务模式 FREE_RESOURCE：释放资源 MODIFY_DB_CONFIG：修改数据库配置 RESET_DB_PWD：重置数据库密码（源库、目标库） MODIFY_CONFIGURATION：修改任务配置 PAUSE：暂停任务 START：启动任务 CHANGE：修改任务 RETRY：重试任务 RESET：重置任务 DELETE：删除任务 QUERY_PRE_CHECK：预检查 SWITCH_OVER：容灾倒换 START_INCR：CASSANDRA启动增量任务 MODIFY_TASK_NUMBER：CASSANDRA修改线程数配置 CONTINUE_JOB：oracle-GaussDB分布式:启动失败或者停止的任务 STOP_JOB：oracle-GaussDB分布式:停止任务 CONTINUE_CAPTURE：oracle-GaussDB分布式:启动抓取 STOP_CAPTURE：oracle-GaussDB分布式:停止抓取 CONTINUE_APPLY：oracle-GaussDB分布式:启动回放 STOP_APPLY：oracle-GaussDB分布式:停止回放 PAY_ORDER：包年包月支付订单 UNSUBSCRIBE：包年包月退订 TO_PERIOD：转包周期 TO_RENEW：包周期续费 ORDER_INFO：订单详情 CHANGE_FLAVOR：规格变更 CLONE：克隆任务
	UnavailableActions *[]JobActionRespUnavailableActions `json:"unavailable_actions,omitempty"`

	// 示例： SWITCH_OVER：灾备倒换中 STOP_JOB：任务暂停中
	CurrentAction *JobActionRespCurrentAction `json:"current_action,omitempty"`
}

func (o JobActionResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobActionResp struct{}"
	}

	return strings.Join([]string{"JobActionResp", string(data)}, " ")
}

type JobActionRespAvailableActions struct {
	value string
}

type JobActionRespAvailableActionsEnum struct {
	CREATE                   JobActionRespAvailableActions
	CHOOSE_OBJECT            JobActionRespAvailableActions
	PRE_CHECK                JobActionRespAvailableActions
	CHANGE_MODE              JobActionRespAvailableActions
	FREE_RESOURCE            JobActionRespAvailableActions
	MODIFY_DB_CONFIG         JobActionRespAvailableActions
	RESET_DB_PWD             JobActionRespAvailableActions
	MODIFY_CONFIGURATION     JobActionRespAvailableActions
	PAUSE                    JobActionRespAvailableActions
	START                    JobActionRespAvailableActions
	CHANGE                   JobActionRespAvailableActions
	RETRY                    JobActionRespAvailableActions
	RESET                    JobActionRespAvailableActions
	DELETE                   JobActionRespAvailableActions
	QUERY_PRE_CHECK          JobActionRespAvailableActions
	SWITCH_OVER              JobActionRespAvailableActions
	MODIFY_SPECIFICATION_ID  JobActionRespAvailableActions
	JUMP_RETRY               JobActionRespAvailableActions
	START_INCR               JobActionRespAvailableActions
	MODIFY_TASK_NUMBER       JobActionRespAvailableActions
	CONTINUE_JOB             JobActionRespAvailableActions
	STOP_JOB                 JobActionRespAvailableActions
	CONTINUE_CAPTURE         JobActionRespAvailableActions
	STOP_CAPTURE             JobActionRespAvailableActions
	CONTINUE_APPLY           JobActionRespAvailableActions
	API_CONFIGURATION_ACTION JobActionRespAvailableActions
	STOP_APPLY               JobActionRespAvailableActions
	PAY_ORDER                JobActionRespAvailableActions
	UNSUBSCRIBE              JobActionRespAvailableActions
	TO_PERIOD                JobActionRespAvailableActions
	TO_RENEW                 JobActionRespAvailableActions
	ORDER_INFO               JobActionRespAvailableActions
	CHANGE_FLAVOR            JobActionRespAvailableActions
	CLONE                    JobActionRespAvailableActions
}

func GetJobActionRespAvailableActionsEnum() JobActionRespAvailableActionsEnum {
	return JobActionRespAvailableActionsEnum{
		CREATE: JobActionRespAvailableActions{
			value: "CREATE",
		},
		CHOOSE_OBJECT: JobActionRespAvailableActions{
			value: "CHOOSE_OBJECT",
		},
		PRE_CHECK: JobActionRespAvailableActions{
			value: "PRE_CHECK",
		},
		CHANGE_MODE: JobActionRespAvailableActions{
			value: "CHANGE_MODE",
		},
		FREE_RESOURCE: JobActionRespAvailableActions{
			value: "FREE_RESOURCE",
		},
		MODIFY_DB_CONFIG: JobActionRespAvailableActions{
			value: "MODIFY_DB_CONFIG",
		},
		RESET_DB_PWD: JobActionRespAvailableActions{
			value: "RESET_DB_PWD",
		},
		MODIFY_CONFIGURATION: JobActionRespAvailableActions{
			value: "MODIFY_CONFIGURATION",
		},
		PAUSE: JobActionRespAvailableActions{
			value: "PAUSE",
		},
		START: JobActionRespAvailableActions{
			value: "START",
		},
		CHANGE: JobActionRespAvailableActions{
			value: "CHANGE",
		},
		RETRY: JobActionRespAvailableActions{
			value: "RETRY",
		},
		RESET: JobActionRespAvailableActions{
			value: "RESET",
		},
		DELETE: JobActionRespAvailableActions{
			value: "DELETE",
		},
		QUERY_PRE_CHECK: JobActionRespAvailableActions{
			value: "QUERY_PRE_CHECK",
		},
		SWITCH_OVER: JobActionRespAvailableActions{
			value: "SWITCH_OVER",
		},
		MODIFY_SPECIFICATION_ID: JobActionRespAvailableActions{
			value: "MODIFY_SPECIFICATION_ID",
		},
		JUMP_RETRY: JobActionRespAvailableActions{
			value: "JUMP_RETRY",
		},
		START_INCR: JobActionRespAvailableActions{
			value: "START_INCR",
		},
		MODIFY_TASK_NUMBER: JobActionRespAvailableActions{
			value: "MODIFY_TASK_NUMBER",
		},
		CONTINUE_JOB: JobActionRespAvailableActions{
			value: "CONTINUE_JOB",
		},
		STOP_JOB: JobActionRespAvailableActions{
			value: "STOP_JOB",
		},
		CONTINUE_CAPTURE: JobActionRespAvailableActions{
			value: "CONTINUE_CAPTURE",
		},
		STOP_CAPTURE: JobActionRespAvailableActions{
			value: "STOP_CAPTURE",
		},
		CONTINUE_APPLY: JobActionRespAvailableActions{
			value: "CONTINUE_APPLY",
		},
		API_CONFIGURATION_ACTION: JobActionRespAvailableActions{
			value: "API_CONFIGURATION_ACTION",
		},
		STOP_APPLY: JobActionRespAvailableActions{
			value: "STOP_APPLY",
		},
		PAY_ORDER: JobActionRespAvailableActions{
			value: "PAY_ORDER",
		},
		UNSUBSCRIBE: JobActionRespAvailableActions{
			value: "UNSUBSCRIBE",
		},
		TO_PERIOD: JobActionRespAvailableActions{
			value: "TO_PERIOD",
		},
		TO_RENEW: JobActionRespAvailableActions{
			value: "TO_RENEW",
		},
		ORDER_INFO: JobActionRespAvailableActions{
			value: "ORDER_INFO",
		},
		CHANGE_FLAVOR: JobActionRespAvailableActions{
			value: "CHANGE_FLAVOR",
		},
		CLONE: JobActionRespAvailableActions{
			value: "CLONE",
		},
	}
}

func (c JobActionRespAvailableActions) Value() string {
	return c.value
}

func (c JobActionRespAvailableActions) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *JobActionRespAvailableActions) UnmarshalJSON(b []byte) error {
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

type JobActionRespUnavailableActions struct {
	value string
}

type JobActionRespUnavailableActionsEnum struct {
	CREATE                   JobActionRespUnavailableActions
	CHOOSE_OBJECT            JobActionRespUnavailableActions
	PRE_CHECK                JobActionRespUnavailableActions
	CHANGE_MODE              JobActionRespUnavailableActions
	FREE_RESOURCE            JobActionRespUnavailableActions
	MODIFY_DB_CONFIG         JobActionRespUnavailableActions
	RESET_DB_PWD             JobActionRespUnavailableActions
	MODIFY_CONFIGURATION     JobActionRespUnavailableActions
	PAUSE                    JobActionRespUnavailableActions
	START                    JobActionRespUnavailableActions
	CHANGE                   JobActionRespUnavailableActions
	RETRY                    JobActionRespUnavailableActions
	RESET                    JobActionRespUnavailableActions
	DELETE                   JobActionRespUnavailableActions
	QUERY_PRE_CHECK          JobActionRespUnavailableActions
	SWITCH_OVER              JobActionRespUnavailableActions
	MODIFY_SPECIFICATION_ID  JobActionRespUnavailableActions
	JUMP_RETRY               JobActionRespUnavailableActions
	START_INCR               JobActionRespUnavailableActions
	MODIFY_TASK_NUMBER       JobActionRespUnavailableActions
	CONTINUE_JOB             JobActionRespUnavailableActions
	STOP_JOB                 JobActionRespUnavailableActions
	CONTINUE_CAPTURE         JobActionRespUnavailableActions
	STOP_CAPTURE             JobActionRespUnavailableActions
	CONTINUE_APPLY           JobActionRespUnavailableActions
	API_CONFIGURATION_ACTION JobActionRespUnavailableActions
	STOP_APPLY               JobActionRespUnavailableActions
	PAY_ORDER                JobActionRespUnavailableActions
	UNSUBSCRIBE              JobActionRespUnavailableActions
	TO_PERIOD                JobActionRespUnavailableActions
	TO_RENEW                 JobActionRespUnavailableActions
	ORDER_INFO               JobActionRespUnavailableActions
	CHANGE_FLAVOR            JobActionRespUnavailableActions
	CLONE                    JobActionRespUnavailableActions
}

func GetJobActionRespUnavailableActionsEnum() JobActionRespUnavailableActionsEnum {
	return JobActionRespUnavailableActionsEnum{
		CREATE: JobActionRespUnavailableActions{
			value: "CREATE",
		},
		CHOOSE_OBJECT: JobActionRespUnavailableActions{
			value: "CHOOSE_OBJECT",
		},
		PRE_CHECK: JobActionRespUnavailableActions{
			value: "PRE_CHECK",
		},
		CHANGE_MODE: JobActionRespUnavailableActions{
			value: "CHANGE_MODE",
		},
		FREE_RESOURCE: JobActionRespUnavailableActions{
			value: "FREE_RESOURCE",
		},
		MODIFY_DB_CONFIG: JobActionRespUnavailableActions{
			value: "MODIFY_DB_CONFIG",
		},
		RESET_DB_PWD: JobActionRespUnavailableActions{
			value: "RESET_DB_PWD",
		},
		MODIFY_CONFIGURATION: JobActionRespUnavailableActions{
			value: "MODIFY_CONFIGURATION",
		},
		PAUSE: JobActionRespUnavailableActions{
			value: "PAUSE",
		},
		START: JobActionRespUnavailableActions{
			value: "START",
		},
		CHANGE: JobActionRespUnavailableActions{
			value: "CHANGE",
		},
		RETRY: JobActionRespUnavailableActions{
			value: "RETRY",
		},
		RESET: JobActionRespUnavailableActions{
			value: "RESET",
		},
		DELETE: JobActionRespUnavailableActions{
			value: "DELETE",
		},
		QUERY_PRE_CHECK: JobActionRespUnavailableActions{
			value: "QUERY_PRE_CHECK",
		},
		SWITCH_OVER: JobActionRespUnavailableActions{
			value: "SWITCH_OVER",
		},
		MODIFY_SPECIFICATION_ID: JobActionRespUnavailableActions{
			value: "MODIFY_SPECIFICATION_ID",
		},
		JUMP_RETRY: JobActionRespUnavailableActions{
			value: "JUMP_RETRY",
		},
		START_INCR: JobActionRespUnavailableActions{
			value: "START_INCR",
		},
		MODIFY_TASK_NUMBER: JobActionRespUnavailableActions{
			value: "MODIFY_TASK_NUMBER",
		},
		CONTINUE_JOB: JobActionRespUnavailableActions{
			value: "CONTINUE_JOB",
		},
		STOP_JOB: JobActionRespUnavailableActions{
			value: "STOP_JOB",
		},
		CONTINUE_CAPTURE: JobActionRespUnavailableActions{
			value: "CONTINUE_CAPTURE",
		},
		STOP_CAPTURE: JobActionRespUnavailableActions{
			value: "STOP_CAPTURE",
		},
		CONTINUE_APPLY: JobActionRespUnavailableActions{
			value: "CONTINUE_APPLY",
		},
		API_CONFIGURATION_ACTION: JobActionRespUnavailableActions{
			value: "API_CONFIGURATION_ACTION",
		},
		STOP_APPLY: JobActionRespUnavailableActions{
			value: "STOP_APPLY",
		},
		PAY_ORDER: JobActionRespUnavailableActions{
			value: "PAY_ORDER",
		},
		UNSUBSCRIBE: JobActionRespUnavailableActions{
			value: "UNSUBSCRIBE",
		},
		TO_PERIOD: JobActionRespUnavailableActions{
			value: "TO_PERIOD",
		},
		TO_RENEW: JobActionRespUnavailableActions{
			value: "TO_RENEW",
		},
		ORDER_INFO: JobActionRespUnavailableActions{
			value: "ORDER_INFO",
		},
		CHANGE_FLAVOR: JobActionRespUnavailableActions{
			value: "CHANGE_FLAVOR",
		},
		CLONE: JobActionRespUnavailableActions{
			value: "CLONE",
		},
	}
}

func (c JobActionRespUnavailableActions) Value() string {
	return c.value
}

func (c JobActionRespUnavailableActions) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *JobActionRespUnavailableActions) UnmarshalJSON(b []byte) error {
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

type JobActionRespCurrentAction struct {
	value string
}

type JobActionRespCurrentActionEnum struct {
	SWITCH_OVER JobActionRespCurrentAction
	STOP_JOB    JobActionRespCurrentAction
}

func GetJobActionRespCurrentActionEnum() JobActionRespCurrentActionEnum {
	return JobActionRespCurrentActionEnum{
		SWITCH_OVER: JobActionRespCurrentAction{
			value: "SWITCH_OVER",
		},
		STOP_JOB: JobActionRespCurrentAction{
			value: "STOP_JOB",
		},
	}
}

func (c JobActionRespCurrentAction) Value() string {
	return c.value
}

func (c JobActionRespCurrentAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *JobActionRespCurrentAction) UnmarshalJSON(b []byte) error {
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
