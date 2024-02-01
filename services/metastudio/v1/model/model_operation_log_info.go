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

	// 命令执行结果。 * USER_CREATE_JOD：用户开始分身数字人定制 * USER_VERIFYING_SUBMITTED：用户提交审核 * SYSTEM_VERIFY_FAILED：自动审核失败 * SYSTEM_VERIFY_SUCCESS：自动审核成功 * ADMIN_VERIFY_SUCCESS：人工审核通过 * ADMIN_VERIFY_FAILED：人工审核不通过 * SYSTEM_TRAIN_DATA_PREPROCESSING：训练数据预处理中 * SYSTEM_TRAIN_DATA_PREPROCESS_FAILED：训练数据预处理失败 * SYSTEM_TRAIN_DATA_PREPROCESS_SUCCESS：训练数据预处理成功 * SYSTEM_TRAINING：开始训练 * ADMIN_STOP_TRAIN：人工终止训练 * SYSTEM_TRAIN_FAILED：训练失败 * SYSTEM_TRAIN_SUCCESS：训练成功 * SYSTEM_INFERENCE_DATA_PREPROCESSING：推理数据预处理中 * SYSTEM_INFERENCE_DATA_PREPROCESS_FAILED：推理数据预处理失败 * SYSTEM_INFERENCE_DATA_PREPROCESS_SUCCESS：推理数据预处理成功 * SYSTEM_JOB_SUCCESS：任务处理完成 * SYSTEM_MARKABLE_VIDEO: 标定视频生成任务 * SYSTEM_MASK_VERIFY_VIDEO: 校验视频生成任务 * SYSTEM_MASK_VERIFY_VIDEO_SUCCESS：校验视频生成成功 * SYSTEM_MASK_VERIFY_VIDEO_FAILED：校验视频生成失败 * SYSTEM_MARKABLE_VIDEO_SUCCESS：标定视频生成成功 * SYSTEM_MARKABLE_VIDEO_FAILED：标定视频生成失败 * SYSTEM_MASK_VIDEO_AND_ACTION_TIME_SUCCESS：自动标定成功 * SYSTEM_MASK_VIDEO_AND_ACTION_TIME_FAILED：自动标定失败 * ADMIN_MASK_UPLOADED：遮罩文件上传完成 * ADMIN_UPDATE_VIDEO：管理员更换视频 * USER_UPDATE_VIDEO：用户更换视频 * ADMIN_MASK_ACTION_TIME：管理员标定
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
	USER_CREATE_JOD                           OperationLogInfoLogType
	USER_VERIFYING_SUBMITTED                  OperationLogInfoLogType
	SYSTEM_VERIFY_FAILED                      OperationLogInfoLogType
	SYSTEM_VERIFY_SUCCESS                     OperationLogInfoLogType
	ADMIN_VERIFY_SUCCESS                      OperationLogInfoLogType
	ADMIN_VERIFY_FAILED                       OperationLogInfoLogType
	SYSTEM_TRAIN_DATA_PREPROCESSING           OperationLogInfoLogType
	SYSTEM_TRAIN_DATA_PREPROCESS_FAILED       OperationLogInfoLogType
	SYSTEM_TRAIN_DATA_PREPROCESS_SUCCESS      OperationLogInfoLogType
	SYSTEM_TRAINING                           OperationLogInfoLogType
	ADMIN_STOP_TRAIN                          OperationLogInfoLogType
	SYSTEM_TRAIN_FAILED                       OperationLogInfoLogType
	SYSTEM_TRAIN_SUCCESS                      OperationLogInfoLogType
	SYSTEM_INFERENCE_DATA_PREPROCESSING       OperationLogInfoLogType
	SYSTEM_INFERENCE_DATA_PREPROCESS_FAILED   OperationLogInfoLogType
	SYSTEM_INFERENCE_DATA_PREPROCESS_SUCCESS  OperationLogInfoLogType
	SYSTEM_JOB_SUCCESS                        OperationLogInfoLogType
	ADMIN_MASK_UPLOADED                       OperationLogInfoLogType
	ADMIN_UPDATE_VIDEO                        OperationLogInfoLogType
	SYSTEM_MARKABLE_VIDEO                     OperationLogInfoLogType
	SYSTEM_MASK_VERIFY_VIDEO                  OperationLogInfoLogType
	SYSTEM_MASK_VERIFY_VIDEO_SUCCESS          OperationLogInfoLogType
	SYSTEM_MASK_VERIFY_VIDEO_FAILED           OperationLogInfoLogType
	SYSTEM_MARKABLE_VIDEO_SUCCESS             OperationLogInfoLogType
	SYSTEM_MARKABLE_VIDEO_FAILED              OperationLogInfoLogType
	SYSTEM_MASK_VIDEO_AND_ACTION_TIME_SUCCESS OperationLogInfoLogType
	SYSTEM_MASK_VIDEO_AND_ACTION_TIME_FAILED  OperationLogInfoLogType
	USER_UPDATE_VIDEO                         OperationLogInfoLogType
	ADMIN_MASK_ACTION_TIME                    OperationLogInfoLogType
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
		SYSTEM_TRAINING: OperationLogInfoLogType{
			value: "SYSTEM_TRAINING",
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
		SYSTEM_MARKABLE_VIDEO_FAILED: OperationLogInfoLogType{
			value: "SYSTEM_MARKABLE_VIDEO_FAILED",
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
		ADMIN_MASK_ACTION_TIME: OperationLogInfoLogType{
			value: "ADMIN_MASK_ACTION_TIME",
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
