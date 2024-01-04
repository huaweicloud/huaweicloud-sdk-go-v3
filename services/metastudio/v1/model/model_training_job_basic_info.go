package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type TrainingJobBasicInfo struct {

	// 任务ID。
	JobId string `json:"job_id"`

	// 分身数字人模型名称。该名称会作为资产库中分身数字人模型资产名称。
	Name string `json:"name"`

	// 任务的状态。 * WAIT_FILE_UPLOAD: 待上传文件 * AUTO_VERIFYING: 自动审核中 * AUTO_VERIFY_FAILED: 自动审核失败 * MANUAL_VERIFYING: 人工审核中 * MANUAL_VERIFY_FAILED: 人工审核失败 * MANUAL_VERIFY_SUCCESS: 审核通过，等待预处理资源 * TRAINING_DATA_PREPROCESSING：训练数据预处理中 * TRAINING_DATA_PREPROCESS_FAILED: 训练数据预处理失败 * TRAINING_DATA_PREPROCESS_SUCCESS: 训练数据预处理完成，等待训练资源中 * TRAINING: 训练中 * TRAIN_FAILED: 训练失败 * TRAIN_SUCCESS: 训练完成，等待预处理资源 * INFERENCE_DATA_PREPROCESSING: 推理数据预处理中 * INFERENCE_DATA_PREPROCESS_FAILED: 推理数据预处理失败 * WAIT_MAIN_FILE_UPLOAD：等待主文件上传 * JOB_SUCCESS: 完成 * WAIT_USER_CONFIRM：等待用户确认训练效果 * JOB_REJECT：驳回 * JOB_PENDING：挂起 * JOB_FINISH：结束，最终状态，不可再做改变
	State TrainingJobBasicInfoState `json:"state"`

	// 模型资产ID。
	AssetId *string `json:"asset_id,omitempty"`

	// 模型资产所属项目ID。
	ProjectId *string `json:"project_id,omitempty"`

	// 分身数字人模型封面下载URL。URL有效期24小时。
	CoverDownloadUrl *string `json:"cover_download_url,omitempty"`

	// 用户最近一次更新任务的时间（包括租户创建或者重新提交），格式遵循：RFC 3339。 例 “2020-07-30T10:43:17Z”
	LastUpdateTime *string `json:"last_update_time,omitempty"`

	// 创建时间，格式遵循：RFC 3339。 例 “2020-07-30T10:43:17Z”
	CreateTime *string `json:"create_time,omitempty"`

	// 分身数字人训练任务创建者联系方式，如手机或邮箱等。
	Contact *string `json:"contact,omitempty"`

	// 分身数字人训练任务的批次名称。
	BatchName *string `json:"batch_name,omitempty"`

	// 分身数字人训练任务标签。
	Tags *[]string `json:"tags,omitempty"`

	// 分身数字人模型版本。默认是V3版本模型。 * V2: V2版本模型 * V3：V3版本模型
	ModelVersion *TrainingJobBasicInfoModelVersion `json:"model_version,omitempty"`

	// 抠图类型。默认是AI。 * AI：AI抠图 * MANUAL：人工抠图
	MattingType *TrainingJobBasicInfoMattingType `json:"matting_type,omitempty"`
}

func (o TrainingJobBasicInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TrainingJobBasicInfo struct{}"
	}

	return strings.Join([]string{"TrainingJobBasicInfo", string(data)}, " ")
}

type TrainingJobBasicInfoState struct {
	value string
}

type TrainingJobBasicInfoStateEnum struct {
	WAIT_FILE_UPLOAD                 TrainingJobBasicInfoState
	AUTO_VERIFYING                   TrainingJobBasicInfoState
	AUTO_VERIFY_FAILED               TrainingJobBasicInfoState
	MANUAL_VERIFYING                 TrainingJobBasicInfoState
	MANUAL_VERIFY_FAILED             TrainingJobBasicInfoState
	MANUAL_VERIFY_SUCCESS            TrainingJobBasicInfoState
	TRAINING_DATA_PREPROCESSING      TrainingJobBasicInfoState
	TRAINING_DATA_PREPROCESS_FAILED  TrainingJobBasicInfoState
	TRAINING_DATA_PREPROCESS_SUCCESS TrainingJobBasicInfoState
	TRAINING                         TrainingJobBasicInfoState
	TRAIN_FAILED                     TrainingJobBasicInfoState
	TRAIN_SUCCESS                    TrainingJobBasicInfoState
	INFERENCE_DATA_PREPROCESSING     TrainingJobBasicInfoState
	INFERENCE_DATA_PREPROCESS_FAILED TrainingJobBasicInfoState
	WAIT_MASK_UPLOAD                 TrainingJobBasicInfoState
	WAIT_MAIN_FILE_UPLOAD            TrainingJobBasicInfoState
	JOB_SUCCESS                      TrainingJobBasicInfoState
	WAIT_USER_CONFIRM                TrainingJobBasicInfoState
	JOB_REJECT                       TrainingJobBasicInfoState
	JOB_PENDING                      TrainingJobBasicInfoState
	JOB_FINISH                       TrainingJobBasicInfoState
}

func GetTrainingJobBasicInfoStateEnum() TrainingJobBasicInfoStateEnum {
	return TrainingJobBasicInfoStateEnum{
		WAIT_FILE_UPLOAD: TrainingJobBasicInfoState{
			value: "WAIT_FILE_UPLOAD",
		},
		AUTO_VERIFYING: TrainingJobBasicInfoState{
			value: "AUTO_VERIFYING",
		},
		AUTO_VERIFY_FAILED: TrainingJobBasicInfoState{
			value: "AUTO_VERIFY_FAILED",
		},
		MANUAL_VERIFYING: TrainingJobBasicInfoState{
			value: "MANUAL_VERIFYING",
		},
		MANUAL_VERIFY_FAILED: TrainingJobBasicInfoState{
			value: "MANUAL_VERIFY_FAILED",
		},
		MANUAL_VERIFY_SUCCESS: TrainingJobBasicInfoState{
			value: "MANUAL_VERIFY_SUCCESS",
		},
		TRAINING_DATA_PREPROCESSING: TrainingJobBasicInfoState{
			value: "TRAINING_DATA_PREPROCESSING",
		},
		TRAINING_DATA_PREPROCESS_FAILED: TrainingJobBasicInfoState{
			value: "TRAINING_DATA_PREPROCESS_FAILED",
		},
		TRAINING_DATA_PREPROCESS_SUCCESS: TrainingJobBasicInfoState{
			value: "TRAINING_DATA_PREPROCESS_SUCCESS",
		},
		TRAINING: TrainingJobBasicInfoState{
			value: "TRAINING",
		},
		TRAIN_FAILED: TrainingJobBasicInfoState{
			value: "TRAIN_FAILED",
		},
		TRAIN_SUCCESS: TrainingJobBasicInfoState{
			value: "TRAIN_SUCCESS",
		},
		INFERENCE_DATA_PREPROCESSING: TrainingJobBasicInfoState{
			value: "INFERENCE_DATA_PREPROCESSING",
		},
		INFERENCE_DATA_PREPROCESS_FAILED: TrainingJobBasicInfoState{
			value: "INFERENCE_DATA_PREPROCESS_FAILED",
		},
		WAIT_MASK_UPLOAD: TrainingJobBasicInfoState{
			value: "WAIT_MASK_UPLOAD",
		},
		WAIT_MAIN_FILE_UPLOAD: TrainingJobBasicInfoState{
			value: "WAIT_MAIN_FILE_UPLOAD",
		},
		JOB_SUCCESS: TrainingJobBasicInfoState{
			value: "JOB_SUCCESS",
		},
		WAIT_USER_CONFIRM: TrainingJobBasicInfoState{
			value: "WAIT_USER_CONFIRM",
		},
		JOB_REJECT: TrainingJobBasicInfoState{
			value: "JOB_REJECT",
		},
		JOB_PENDING: TrainingJobBasicInfoState{
			value: "JOB_PENDING",
		},
		JOB_FINISH: TrainingJobBasicInfoState{
			value: "JOB_FINISH",
		},
	}
}

func (c TrainingJobBasicInfoState) Value() string {
	return c.value
}

func (c TrainingJobBasicInfoState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TrainingJobBasicInfoState) UnmarshalJSON(b []byte) error {
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

type TrainingJobBasicInfoModelVersion struct {
	value string
}

type TrainingJobBasicInfoModelVersionEnum struct {
	V2 TrainingJobBasicInfoModelVersion
	V3 TrainingJobBasicInfoModelVersion
}

func GetTrainingJobBasicInfoModelVersionEnum() TrainingJobBasicInfoModelVersionEnum {
	return TrainingJobBasicInfoModelVersionEnum{
		V2: TrainingJobBasicInfoModelVersion{
			value: "V2",
		},
		V3: TrainingJobBasicInfoModelVersion{
			value: "V3",
		},
	}
}

func (c TrainingJobBasicInfoModelVersion) Value() string {
	return c.value
}

func (c TrainingJobBasicInfoModelVersion) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TrainingJobBasicInfoModelVersion) UnmarshalJSON(b []byte) error {
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

type TrainingJobBasicInfoMattingType struct {
	value string
}

type TrainingJobBasicInfoMattingTypeEnum struct {
	AI     TrainingJobBasicInfoMattingType
	MANUAL TrainingJobBasicInfoMattingType
}

func GetTrainingJobBasicInfoMattingTypeEnum() TrainingJobBasicInfoMattingTypeEnum {
	return TrainingJobBasicInfoMattingTypeEnum{
		AI: TrainingJobBasicInfoMattingType{
			value: "AI",
		},
		MANUAL: TrainingJobBasicInfoMattingType{
			value: "MANUAL",
		},
	}
}

func (c TrainingJobBasicInfoMattingType) Value() string {
	return c.value
}

func (c TrainingJobBasicInfoMattingType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TrainingJobBasicInfoMattingType) UnmarshalJSON(b []byte) error {
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
