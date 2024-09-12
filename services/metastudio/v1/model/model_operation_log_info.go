package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// OperationLogInfo 操作日志。
type OperationLogInfo struct {

	// 操作时间,格式遵循：RFC 3339。 例 “2020-07-30T10:43:17Z”
	OperateTime *string `json:"operate_time,omitempty"`

	// 命令执行结果。 * USER_CREATE_JOD：用户开始分身数字人定制 * USER_VERIFYING_SUBMITTED：用户提交审核 * SYSTEM_VERIFY_FAILED：自动审核失败 * SYSTEM_VERIFY_SUCCESS：自动审核成功 * ADMIN_VERIFY_SUCCESS：人工审核通过 * ADMIN_VERIFY_FAILED：人工审核不通过 * SYSTEM_TRAIN_DATA_PREPROCESSING：训练数据预处理中 * SYSTEM_TRAIN_DATA_PREPROCESS_FAILED：训练数据预处理失 * SYSTEM_TRAIN_DATA_PREPROCESS_SUCCESS：训练数据预处理成功 * SYSTEM_ACTION_GENERATE_DATA_PREPROCESSING：动作编排原子动作生成中 * SYSTEM_ACTION_GENERATE_DATA_SUCCESS：动作编排原子动作生成成功 * SYSTEM_ACTION_GENERATE_ORI_SUCCESS：动作编排资产数据生成成功 * SYSTEM_ACTION_GENERATE_DATA_FAILED：动作编排原子动作生成失败 * SYSTEM_ACTION_GENERATE_ORI_FAILED：动作编排资产数据生成失败 * SYSTEM_ACTION_GENERATE_ORI_PREPROCESSING：动作编排资产数据生成中 * SYSTEM_TRAINING：开始训练 * ADMIN_STOP_TRAINING_DATA_PREPROCESS：人工中止训练 * ADMIN_STOP_INFERENCE_DATA_PREPROCESS：人工中止推理预处理 * ADMIN_STOP_TRAIN：人工终止训练 * SYSTEM_TRAIN_FAILED：训练失败 * SYSTEM_TRAIN_SUCCESS：训练成功 * SYSTEM_INFERENCE_DATA_PREPROCESSING：推理数据预处理中 * SYSTEM_INFERENCE_DATA_PREPROCESS_FAILED：推理数据预处理失败 * SYSTEM_INFERENCE_DATA_PREPROCESS_SUCCESS：推理数据预处理成功 * SYSTEM_JOB_SUCCESS：任务处理完成 * ADMIN_MASK_UPLOADED：遮罩文件上传完成（已废弃） * ADMIN_UPDATE_VIDEO：管理员更换视频 * ADMIN_UPDATE_ACTION_VIDEO：管理员更换动作编排视频 * ADMIN_RESET：管理员一键重置 * ADMIN_ACCEPT：管理员通过 * USER_REPAIR：用户修复 * SYSTEM_UPDATE_COVER：更换封面 * ADMIN_SET_SILENCE_REPEAT_NUM：管理员设置静默轮数 * SYSTEM_MARKABLE_VIDEO：标记视频生成任务 * SYSTEM_MASK_VERIFY_VIDEO：校验视频生成任务 * SYSTEM_MASK_VERIFY_VIDEO_SUCCESS：校验视频生成成功 * SYSTEM_MASK_VERIFY_VIDEO_FAILED：校验视频生成失败 * SYSTEM_MARKABLE_VIDEO_SUCCESS：标记视频生成成功 * ADMIN_CONFIRM_ACTION：管理员确认动作 * ADMIN_STOP_ACTION_GENERATE_DATA_PREPROCESS：人工中止原子动作生成 * ADMIN_STOP_ACTION_GENERATE_ORI_PREPROCESS：人工中止动作编排 * SYSTEM_COMPILE_FAILED：转编译失败 * SYSTEM_COMPILE_SUCCESS：转编译成功 * SYSTEM_MARKABLE_VIDEO_FAILED：标记视频生成失败 * ADMIN_UPDATE_COMPILE：管理员更新转编译配置 * ADMIN_UPDATE_INFERENCE_DATA_PROCESS_VIDEO：管理员更新推理预处里时间段信息 * SYSTEM_EXECUTE_COMPILE：执行转编译 * SYSTEM_MASK_VIDEO_AND_ACTION_TIME_SUCCESS：自动标记成功 * SYSTEM_MASK_VIDEO_AND_ACTION_TIME_FAILED：自动标记失败 * USER_UPDATE_VIDEO：用户更换视频 * ADMIN_UPDATE_GENERAL_CONFIG：管理员更新通用配置 * ADMIN_MASK_ACTION_TIME：管理员标记 * STOP_COMPILE：人工中止转编译
	LogType *OperationLogInfoLogType `json:"log_type,omitempty"`

	// 日志描述。用于记录人工审核不通过时的审核意见和DHTS、DHPS上报的错误信息。
	LogDescription *string `json:"log_description,omitempty"`

	// 操作人员。 * USER：用户 * ADMIN：管理员 * SYSTEM：系统
	OperateUser *OperationLogInfoOperateUser `json:"operate_user,omitempty"`
}

func (o OperationLogInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OperationLogInfo struct{}"
	}

	return strings.Join([]string{"OperationLogInfo", string(data)}, " ")
}

type OperationLogInfoLogType struct {
	value string
}

type OperationLogInfoLogTypeEnum struct {
	USER_CREATE_JOD                            OperationLogInfoLogType
	USER_VERIFYING_SUBMITTED                   OperationLogInfoLogType
	SYSTEM_VERIFY_FAILED                       OperationLogInfoLogType
	ADMIN_UPDATE_BEAUTY_LEVEL                  OperationLogInfoLogType
	SYSTEM_VERIFY_SUCCESS                      OperationLogInfoLogType
	ADMIN_VERIFY_SUCCESS                       OperationLogInfoLogType
	ADMIN_VERIFY_FAILED                        OperationLogInfoLogType
	SYSTEM_TRAIN_DATA_PREPROCESSING            OperationLogInfoLogType
	SYSTEM_TRAIN_DATA_PREPROCESS_FAILED        OperationLogInfoLogType
	SYSTEM_TRAIN_DATA_PREPROCESS_SUCCESS       OperationLogInfoLogType
	SYSTEM_ACTION_GENERATE_DATA_PREPROCESSING  OperationLogInfoLogType
	SYSTEM_ACTION_GENERATE_DATA_SUCCESS        OperationLogInfoLogType
	SYSTEM_ACTION_GENERATE_ORI_SUCCESS         OperationLogInfoLogType
	SYSTEM_ACTION_GENERATE_DATA_FAILED         OperationLogInfoLogType
	SYSTEM_ACTION_GENERATE_ORI_FAILED          OperationLogInfoLogType
	SYSTEM_ACTION_GENERATE_ORI_PREPROCESSING   OperationLogInfoLogType
	SYSTEM_TRAINING                            OperationLogInfoLogType
	ADMIN_STOP_TRAINING_DATA_PREPROCESS        OperationLogInfoLogType
	ADMIN_STOP_BEAUTY_PREPROCESS               OperationLogInfoLogType
	ADMIN_STOP_INFERENCE_DATA_PREPROCESS       OperationLogInfoLogType
	ADMIN_STOP_TRAIN                           OperationLogInfoLogType
	SYSTEM_TRAIN_FAILED                        OperationLogInfoLogType
	SYSTEM_TRAIN_SUCCESS                       OperationLogInfoLogType
	SYSTEM_INFERENCE_DATA_PREPROCESSING        OperationLogInfoLogType
	SYSTEM_INFERENCE_DATA_PREPROCESS_FAILED    OperationLogInfoLogType
	SYSTEM_INFERENCE_DATA_PREPROCESS_SUCCESS   OperationLogInfoLogType
	SYSTEM_JOB_SUCCESS                         OperationLogInfoLogType
	ADMIN_MASK_UPLOADED                        OperationLogInfoLogType
	ADMIN_UPDATE_VIDEO                         OperationLogInfoLogType
	ADMIN_UPDATE_ACTION_VIDEO                  OperationLogInfoLogType
	ADMIN_RESET                                OperationLogInfoLogType
	ADMIN_ACCEPT                               OperationLogInfoLogType
	USER_REPAIR                                OperationLogInfoLogType
	SYSTEM_UPDATE_COVER                        OperationLogInfoLogType
	ADMIN_SET_SILENCE_REPEAT_NUM               OperationLogInfoLogType
	SYSTEM_MARKABLE_VIDEO                      OperationLogInfoLogType
	SYSTEM_MASK_VERIFY_VIDEO                   OperationLogInfoLogType
	SYSTEM_MASK_VERIFY_VIDEO_SUCCESS           OperationLogInfoLogType
	SYSTEM_MASK_VERIFY_VIDEO_FAILED            OperationLogInfoLogType
	SYSTEM_MARKABLE_VIDEO_SUCCESS              OperationLogInfoLogType
	SYSTEM_BEAUTY_PREPROCESSING                OperationLogInfoLogType
	SYSTEM_BEAUTY_PREPROCESS_FAILED            OperationLogInfoLogType
	ADMIN_CONFIRM_ACTION                       OperationLogInfoLogType
	ADMIN_STOP_ACTION_GENERATE_DATA_PREPROCESS OperationLogInfoLogType
	ADMIN_STOP_ACTION_GENERATE_ORI_PREPROCESS  OperationLogInfoLogType
	SYSTEM_BEAUTY_PREPROCESS_SUCCESS           OperationLogInfoLogType
	SYSTEM_COMPILE_FAILED                      OperationLogInfoLogType
	SYSTEM_COMPILE_SUCCESS                     OperationLogInfoLogType
	SYSTEM_MARKABLE_VIDEO_FAILED               OperationLogInfoLogType
	ADMIN_UPDATE_COMPILE                       OperationLogInfoLogType
	ADMIN_UPDATE_INFERENCE_DATA_PROCESS_VIDEO  OperationLogInfoLogType
	SYSTEM_EXECUTE_COMPILE                     OperationLogInfoLogType
	SYSTEM_EXECUTE_BEAUTY                      OperationLogInfoLogType
	SYSTEM_MASK_VIDEO_AND_ACTION_TIME_SUCCESS  OperationLogInfoLogType
	SYSTEM_MASK_VIDEO_AND_ACTION_TIME_FAILED   OperationLogInfoLogType
	USER_UPDATE_VIDEO                          OperationLogInfoLogType
	ADMIN_UPDATE_GENERAL_CONFIG                OperationLogInfoLogType
	ADMIN_MASK_ACTION_TIME                     OperationLogInfoLogType
	STOP_COMPILE                               OperationLogInfoLogType
}

func GetOperationLogInfoLogTypeEnum() OperationLogInfoLogTypeEnum {
	return OperationLogInfoLogTypeEnum{
		USER_CREATE_JOD: OperationLogInfoLogType{
			value: "USER_CREATE_JOD",
		},
		USER_VERIFYING_SUBMITTED: OperationLogInfoLogType{
			value: "USER_VERIFYING_SUBMITTED",
		},
		SYSTEM_VERIFY_FAILED: OperationLogInfoLogType{
			value: "SYSTEM_VERIFY_FAILED",
		},
		ADMIN_UPDATE_BEAUTY_LEVEL: OperationLogInfoLogType{
			value: "ADMIN_UPDATE_BEAUTY_LEVEL",
		},
		SYSTEM_VERIFY_SUCCESS: OperationLogInfoLogType{
			value: "SYSTEM_VERIFY_SUCCESS",
		},
		ADMIN_VERIFY_SUCCESS: OperationLogInfoLogType{
			value: "ADMIN_VERIFY_SUCCESS",
		},
		ADMIN_VERIFY_FAILED: OperationLogInfoLogType{
			value: "ADMIN_VERIFY_FAILED",
		},
		SYSTEM_TRAIN_DATA_PREPROCESSING: OperationLogInfoLogType{
			value: "SYSTEM_TRAIN_DATA_PREPROCESSING",
		},
		SYSTEM_TRAIN_DATA_PREPROCESS_FAILED: OperationLogInfoLogType{
			value: "SYSTEM_TRAIN_DATA_PREPROCESS_FAILED",
		},
		SYSTEM_TRAIN_DATA_PREPROCESS_SUCCESS: OperationLogInfoLogType{
			value: "SYSTEM_TRAIN_DATA_PREPROCESS_SUCCESS",
		},
		SYSTEM_ACTION_GENERATE_DATA_PREPROCESSING: OperationLogInfoLogType{
			value: "SYSTEM_ACTION_GENERATE_DATA_PREPROCESSING",
		},
		SYSTEM_ACTION_GENERATE_DATA_SUCCESS: OperationLogInfoLogType{
			value: "SYSTEM_ACTION_GENERATE_DATA_SUCCESS",
		},
		SYSTEM_ACTION_GENERATE_ORI_SUCCESS: OperationLogInfoLogType{
			value: "SYSTEM_ACTION_GENERATE_ORI_SUCCESS",
		},
		SYSTEM_ACTION_GENERATE_DATA_FAILED: OperationLogInfoLogType{
			value: "SYSTEM_ACTION_GENERATE_DATA_FAILED",
		},
		SYSTEM_ACTION_GENERATE_ORI_FAILED: OperationLogInfoLogType{
			value: "SYSTEM_ACTION_GENERATE_ORI_FAILED",
		},
		SYSTEM_ACTION_GENERATE_ORI_PREPROCESSING: OperationLogInfoLogType{
			value: "SYSTEM_ACTION_GENERATE_ORI_PREPROCESSING",
		},
		SYSTEM_TRAINING: OperationLogInfoLogType{
			value: "SYSTEM_TRAINING",
		},
		ADMIN_STOP_TRAINING_DATA_PREPROCESS: OperationLogInfoLogType{
			value: "ADMIN_STOP_TRAINING_DATA_PREPROCESS",
		},
		ADMIN_STOP_BEAUTY_PREPROCESS: OperationLogInfoLogType{
			value: "ADMIN_STOP_BEAUTY_PREPROCESS",
		},
		ADMIN_STOP_INFERENCE_DATA_PREPROCESS: OperationLogInfoLogType{
			value: "ADMIN_STOP_INFERENCE_DATA_PREPROCESS",
		},
		ADMIN_STOP_TRAIN: OperationLogInfoLogType{
			value: "ADMIN_STOP_TRAIN",
		},
		SYSTEM_TRAIN_FAILED: OperationLogInfoLogType{
			value: "SYSTEM_TRAIN_FAILED",
		},
		SYSTEM_TRAIN_SUCCESS: OperationLogInfoLogType{
			value: "SYSTEM_TRAIN_SUCCESS",
		},
		SYSTEM_INFERENCE_DATA_PREPROCESSING: OperationLogInfoLogType{
			value: "SYSTEM_INFERENCE_DATA_PREPROCESSING",
		},
		SYSTEM_INFERENCE_DATA_PREPROCESS_FAILED: OperationLogInfoLogType{
			value: "SYSTEM_INFERENCE_DATA_PREPROCESS_FAILED",
		},
		SYSTEM_INFERENCE_DATA_PREPROCESS_SUCCESS: OperationLogInfoLogType{
			value: "SYSTEM_INFERENCE_DATA_PREPROCESS_SUCCESS",
		},
		SYSTEM_JOB_SUCCESS: OperationLogInfoLogType{
			value: "SYSTEM_JOB_SUCCESS",
		},
		ADMIN_MASK_UPLOADED: OperationLogInfoLogType{
			value: "ADMIN_MASK_UPLOADED",
		},
		ADMIN_UPDATE_VIDEO: OperationLogInfoLogType{
			value: "ADMIN_UPDATE_VIDEO",
		},
		ADMIN_UPDATE_ACTION_VIDEO: OperationLogInfoLogType{
			value: "ADMIN_UPDATE_ACTION_VIDEO",
		},
		ADMIN_RESET: OperationLogInfoLogType{
			value: "ADMIN_RESET",
		},
		ADMIN_ACCEPT: OperationLogInfoLogType{
			value: "ADMIN_ACCEPT",
		},
		USER_REPAIR: OperationLogInfoLogType{
			value: "USER_REPAIR",
		},
		SYSTEM_UPDATE_COVER: OperationLogInfoLogType{
			value: "SYSTEM_UPDATE_COVER",
		},
		ADMIN_SET_SILENCE_REPEAT_NUM: OperationLogInfoLogType{
			value: "ADMIN_SET_SILENCE_REPEAT_NUM",
		},
		SYSTEM_MARKABLE_VIDEO: OperationLogInfoLogType{
			value: "SYSTEM_MARKABLE_VIDEO",
		},
		SYSTEM_MASK_VERIFY_VIDEO: OperationLogInfoLogType{
			value: "SYSTEM_MASK_VERIFY_VIDEO",
		},
		SYSTEM_MASK_VERIFY_VIDEO_SUCCESS: OperationLogInfoLogType{
			value: "SYSTEM_MASK_VERIFY_VIDEO_SUCCESS",
		},
		SYSTEM_MASK_VERIFY_VIDEO_FAILED: OperationLogInfoLogType{
			value: "SYSTEM_MASK_VERIFY_VIDEO_FAILED",
		},
		SYSTEM_MARKABLE_VIDEO_SUCCESS: OperationLogInfoLogType{
			value: "SYSTEM_MARKABLE_VIDEO_SUCCESS",
		},
		SYSTEM_BEAUTY_PREPROCESSING: OperationLogInfoLogType{
			value: "SYSTEM_BEAUTY_PREPROCESSING",
		},
		SYSTEM_BEAUTY_PREPROCESS_FAILED: OperationLogInfoLogType{
			value: "SYSTEM_BEAUTY_PREPROCESS_FAILED",
		},
		ADMIN_CONFIRM_ACTION: OperationLogInfoLogType{
			value: "ADMIN_CONFIRM_ACTION",
		},
		ADMIN_STOP_ACTION_GENERATE_DATA_PREPROCESS: OperationLogInfoLogType{
			value: "ADMIN_STOP_ACTION_GENERATE_DATA_PREPROCESS",
		},
		ADMIN_STOP_ACTION_GENERATE_ORI_PREPROCESS: OperationLogInfoLogType{
			value: "ADMIN_STOP_ACTION_GENERATE_ORI_PREPROCESS",
		},
		SYSTEM_BEAUTY_PREPROCESS_SUCCESS: OperationLogInfoLogType{
			value: "SYSTEM_BEAUTY_PREPROCESS_SUCCESS",
		},
		SYSTEM_COMPILE_FAILED: OperationLogInfoLogType{
			value: "SYSTEM_COMPILE_FAILED",
		},
		SYSTEM_COMPILE_SUCCESS: OperationLogInfoLogType{
			value: "SYSTEM_COMPILE_SUCCESS",
		},
		SYSTEM_MARKABLE_VIDEO_FAILED: OperationLogInfoLogType{
			value: "SYSTEM_MARKABLE_VIDEO_FAILED",
		},
		ADMIN_UPDATE_COMPILE: OperationLogInfoLogType{
			value: "ADMIN_UPDATE_COMPILE",
		},
		ADMIN_UPDATE_INFERENCE_DATA_PROCESS_VIDEO: OperationLogInfoLogType{
			value: "ADMIN_UPDATE_INFERENCE_DATA_PROCESS_VIDEO",
		},
		SYSTEM_EXECUTE_COMPILE: OperationLogInfoLogType{
			value: "SYSTEM_EXECUTE_COMPILE",
		},
		SYSTEM_EXECUTE_BEAUTY: OperationLogInfoLogType{
			value: "SYSTEM_EXECUTE_BEAUTY",
		},
		SYSTEM_MASK_VIDEO_AND_ACTION_TIME_SUCCESS: OperationLogInfoLogType{
			value: "SYSTEM_MASK_VIDEO_AND_ACTION_TIME_SUCCESS",
		},
		SYSTEM_MASK_VIDEO_AND_ACTION_TIME_FAILED: OperationLogInfoLogType{
			value: "SYSTEM_MASK_VIDEO_AND_ACTION_TIME_FAILED",
		},
		USER_UPDATE_VIDEO: OperationLogInfoLogType{
			value: "USER_UPDATE_VIDEO",
		},
		ADMIN_UPDATE_GENERAL_CONFIG: OperationLogInfoLogType{
			value: "ADMIN_UPDATE_GENERAL_CONFIG",
		},
		ADMIN_MASK_ACTION_TIME: OperationLogInfoLogType{
			value: "ADMIN_MASK_ACTION_TIME",
		},
		STOP_COMPILE: OperationLogInfoLogType{
			value: "STOP_COMPILE",
		},
	}
}

func (c OperationLogInfoLogType) Value() string {
	return c.value
}

func (c OperationLogInfoLogType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *OperationLogInfoLogType) UnmarshalJSON(b []byte) error {
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

type OperationLogInfoOperateUser struct {
	value string
}

type OperationLogInfoOperateUserEnum struct {
	USER   OperationLogInfoOperateUser
	ADMIN  OperationLogInfoOperateUser
	SYSTEM OperationLogInfoOperateUser
}

func GetOperationLogInfoOperateUserEnum() OperationLogInfoOperateUserEnum {
	return OperationLogInfoOperateUserEnum{
		USER: OperationLogInfoOperateUser{
			value: "USER",
		},
		ADMIN: OperationLogInfoOperateUser{
			value: "ADMIN",
		},
		SYSTEM: OperationLogInfoOperateUser{
			value: "SYSTEM",
		},
	}
}

func (c OperationLogInfoOperateUser) Value() string {
	return c.value
}

func (c OperationLogInfoOperateUser) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *OperationLogInfoOperateUser) UnmarshalJSON(b []byte) error {
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
