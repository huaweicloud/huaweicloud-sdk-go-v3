package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// JobActions 任务操作命令集合。
type JobActions struct {

	// 任务可操作命令集合。 - CREATE：创建任务 - CHOOSE_OBJECT：选择对象，任务增量中再编辑 - PRE_CHECK：预检查 - CHANGE_MODE：修改任务模式 - FREE_RESOURCE：释放资源 - MODIFY_DB_CONFIG：修改数据库配置 - RESET_DB_PWD：重置数据库密码（源库、目标库） - MODIFY_CONFIGURATION：修改任务配置 - PAUSE：暂停任务 - START：启动任务 - CHANGE：修改任务 - RETRY：重试任务 - RESET：重置任务 - DELETE：删除任务 - QUERY_PRE_CHECK：预检查 - SWITCH_OVER：容灾倒换 - START_INCR：CASSANDRA启动增量任务 - MODIFY_TASK_NUMBER：CASSANDRA修改线程数配置 - CONTINUE_JOB：oracle-GaussDB分布式:启动失败或者停止的任务 - STOP_JOB：oracle-GaussDB分布式:停止任务 - CONTINUE_CAPTURE：启动抓取 - STOP_CAPTURE：停止抓取 - CONTINUE_APPLY：启动回放 - STOP_APPLY：停止回放 - PAY_ORDER：包年包月支付订单 - UNSUBSCRIBE：包年包月退订 - TO_PERIOD：转包周期 - TO_RENEW：包周期续费 - ORDER_INFO：订单详情 - CHANGE_FLAVOR：规格变更 - CLONE：克隆任务
	AvailableActions []string `json:"available_actions"`

	// 任务可操作命令集合。 - CREATE：创建任务 - CHOOSE_OBJECT：选择对象，任务增量中再编辑 - PRE_CHECK：预检查 - CHANGE_MODE：修改任务模式 - FREE_RESOURCE：释放资源 - MODIFY_DB_CONFIG：修改数据库配置 - RESET_DB_PWD：重置数据库密码（源库、目标库） - MODIFY_CONFIGURATION：修改任务配置 - PAUSE：暂停任务 - START：启动任务 - CHANGE：修改任务 - RETRY：重试任务 - RESET：重置任务 - DELETE：删除任务 - QUERY_PRE_CHECK：预检查 - SWITCH_OVER：容灾倒换 - START_INCR：CASSANDRA启动增量任务 - MODIFY_TASK_NUMBER：CASSANDRA修改线程数配置 - CONTINUE_JOB：oracle-GaussDB分布式:启动失败或者停止的任务 - STOP_JOB：oracle-GaussDB分布式:停止任务 - CONTINUE_CAPTURE：启动抓取 - STOP_CAPTURE：停止抓取 - CONTINUE_APPLY：启动回放 - STOP_APPLY：停止回放 - PAY_ORDER：包年包月支付订单 - UNSUBSCRIBE：包年包月退订 - TO_PERIOD：转包周期 - TO_RENEW：包周期续费 - ORDER_INFO：订单详情 - CHANGE_FLAVOR：规格变更 - CLONE：克隆任务
	UnavailableActions []string `json:"unavailable_actions"`

	// 任务当前操作命令。取值： - API_CONFIGURATION_ACTION：OPEN API配置中的任务能调用。 - CHANGE：修改任务。 - CHANGE_MODE：修改任务模式。 - CHOOSE_OBJECT：选择对象。 - CLONE：克隆任务。 - CONTINUE_APPLY：启动回放，Oracle同步到GaussDB分布式适用。 - CONTINUE_CAPTURE：启动抓取，Oracle同步到GaussDB分布式适用。 - CONTINUE_JOB：启动失败或者停止的任务，Oracle同步到GaussDB分布式适用。 - CREATE：创建任务。 - DELETE：删除任务。 - FREE_RESOURCE：释放资源。 - JUMP_RETRY：跳跃续传任务。 - MODIFY_CONFIGURATION：修改任务配置。 - MODIFY_DB_CONFIG：修改数据库配置。 - MODIFY_TASK_NUMBER：修改线程数配置。 - NODE_FLAVOR_MODIFY：规格变更。 - ORDER_INFO：订单详情。 - PAUSE：暂停任务。 - PAY_ORDER：包年/包月支付订单。 - PRE_CHECK：预检查。 - QUERY_PRE_CHECK：查询预检查结果。 - RESET：重置任务。 - RESET_DB_PWD：重置数据库密码（源库、目标库）。 - RETRY：重试任务。 - START：启动任务。 - START_INCR：启动增量任务。 - STOP_APPLY：停止回放，Oracle同步到GaussDB分布式适用。 - STOP_CAPTURE：停止抓取，Oracle同步到GaussDB分布式适用。 - STOP_JOB：停止任务，Oracle同步到GaussDB分布式适用。 - SWITCH_OVER：灾备倒换。 - TO_PERIOD：转包年/包月任务。 - TO_RENEW：包年/包月任务续费。 - UNSUBSCRIBE：包年/包月任务退订。
	CurrentAction JobActionsCurrentAction `json:"current_action"`
}

func (o JobActions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobActions struct{}"
	}

	return strings.Join([]string{"JobActions", string(data)}, " ")
}

type JobActionsCurrentAction struct {
	value string
}

type JobActionsCurrentActionEnum struct {
	API_CONFIGURATION_ACTION JobActionsCurrentAction
	CHANGE                   JobActionsCurrentAction
	CHANGE_MODE              JobActionsCurrentAction
	CHOOSE_OBJECT            JobActionsCurrentAction
	CLONE                    JobActionsCurrentAction
	CONTINUE_APPLY           JobActionsCurrentAction
	CONTINUE_CAPTURE         JobActionsCurrentAction
	CONTINUE_JOB             JobActionsCurrentAction
	CREATE                   JobActionsCurrentAction
	DELETE                   JobActionsCurrentAction
	FREE_RESOURCE            JobActionsCurrentAction
	JUMP_RETRY               JobActionsCurrentAction
	MODIFY_CONFIGURATION     JobActionsCurrentAction
	MODIFY_DB_CONFIG         JobActionsCurrentAction
	MODIFY_TASK_NUMBER       JobActionsCurrentAction
	NODE_FLAVOR_MODIFY       JobActionsCurrentAction
	ORDER_INFO               JobActionsCurrentAction
	PAUSE                    JobActionsCurrentAction
	PAY_ORDER                JobActionsCurrentAction
	PRE_CHECK                JobActionsCurrentAction
	QUERY_PRE_CHECK          JobActionsCurrentAction
	RESET                    JobActionsCurrentAction
	RESET_DB_PWD             JobActionsCurrentAction
	RETRY                    JobActionsCurrentAction
	START                    JobActionsCurrentAction
	START_INCR               JobActionsCurrentAction
	STOP_APPLY               JobActionsCurrentAction
	STOP_CAPTURE             JobActionsCurrentAction
	STOP_JOB                 JobActionsCurrentAction
	SWITCH_OVER              JobActionsCurrentAction
	TO_PERIOD                JobActionsCurrentAction
	TO_RENEW                 JobActionsCurrentAction
	UNSUBSCRIBE              JobActionsCurrentAction
}

func GetJobActionsCurrentActionEnum() JobActionsCurrentActionEnum {
	return JobActionsCurrentActionEnum{
		API_CONFIGURATION_ACTION: JobActionsCurrentAction{
			value: "API_CONFIGURATION_ACTION",
		},
		CHANGE: JobActionsCurrentAction{
			value: "CHANGE",
		},
		CHANGE_MODE: JobActionsCurrentAction{
			value: "CHANGE_MODE",
		},
		CHOOSE_OBJECT: JobActionsCurrentAction{
			value: "CHOOSE_OBJECT",
		},
		CLONE: JobActionsCurrentAction{
			value: "CLONE",
		},
		CONTINUE_APPLY: JobActionsCurrentAction{
			value: "CONTINUE_APPLY",
		},
		CONTINUE_CAPTURE: JobActionsCurrentAction{
			value: "CONTINUE_CAPTURE",
		},
		CONTINUE_JOB: JobActionsCurrentAction{
			value: "CONTINUE_JOB",
		},
		CREATE: JobActionsCurrentAction{
			value: "CREATE",
		},
		DELETE: JobActionsCurrentAction{
			value: "DELETE",
		},
		FREE_RESOURCE: JobActionsCurrentAction{
			value: "FREE_RESOURCE",
		},
		JUMP_RETRY: JobActionsCurrentAction{
			value: "JUMP_RETRY",
		},
		MODIFY_CONFIGURATION: JobActionsCurrentAction{
			value: "MODIFY_CONFIGURATION",
		},
		MODIFY_DB_CONFIG: JobActionsCurrentAction{
			value: "MODIFY_DB_CONFIG",
		},
		MODIFY_TASK_NUMBER: JobActionsCurrentAction{
			value: "MODIFY_TASK_NUMBER",
		},
		NODE_FLAVOR_MODIFY: JobActionsCurrentAction{
			value: "NODE_FLAVOR_MODIFY",
		},
		ORDER_INFO: JobActionsCurrentAction{
			value: "ORDER_INFO",
		},
		PAUSE: JobActionsCurrentAction{
			value: "PAUSE",
		},
		PAY_ORDER: JobActionsCurrentAction{
			value: "PAY_ORDER",
		},
		PRE_CHECK: JobActionsCurrentAction{
			value: "PRE_CHECK",
		},
		QUERY_PRE_CHECK: JobActionsCurrentAction{
			value: "QUERY_PRE_CHECK",
		},
		RESET: JobActionsCurrentAction{
			value: "RESET",
		},
		RESET_DB_PWD: JobActionsCurrentAction{
			value: "RESET_DB_PWD",
		},
		RETRY: JobActionsCurrentAction{
			value: "RETRY",
		},
		START: JobActionsCurrentAction{
			value: "START",
		},
		START_INCR: JobActionsCurrentAction{
			value: "START_INCR",
		},
		STOP_APPLY: JobActionsCurrentAction{
			value: "STOP_APPLY",
		},
		STOP_CAPTURE: JobActionsCurrentAction{
			value: "STOP_CAPTURE",
		},
		STOP_JOB: JobActionsCurrentAction{
			value: "STOP_JOB",
		},
		SWITCH_OVER: JobActionsCurrentAction{
			value: "SWITCH_OVER",
		},
		TO_PERIOD: JobActionsCurrentAction{
			value: "TO_PERIOD",
		},
		TO_RENEW: JobActionsCurrentAction{
			value: "TO_RENEW",
		},
		UNSUBSCRIBE: JobActionsCurrentAction{
			value: "UNSUBSCRIBE",
		},
	}
}

func (c JobActionsCurrentAction) Value() string {
	return c.value
}

func (c JobActionsCurrentAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *JobActionsCurrentAction) UnmarshalJSON(b []byte) error {
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
