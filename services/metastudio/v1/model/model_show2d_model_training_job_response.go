package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Show2dModelTrainingJobResponse Response Object
type Show2dModelTrainingJobResponse struct {

	// 任务ID。
	JobId string `json:"job_id"`

	// 分身数字人模型名称。该名称会作为资产库中分身数字人模型资产名称。
	Name string `json:"name"`

	// 任务的状态。 * WAIT_FILE_UPLOAD: 待上传文件 * AUTO_VERIFYING: 自动审核中 * AUTO_VERIFY_FAILED: 自动审核失败 * MANUAL_VERIFYING: 人工审核中 * MANUAL_VERIFY_FAILED: 人工审核失败 * MANUAL_VERIFY_SUCCESS: 审核通过，等待预处理资源 * TRAINING_DATA_PREPROCESSING：训练数据预处理中 * TRAINING_DATA_PREPROCESS_FAILED: 训练数据预处理失败 * TRAINING_DATA_PREPROCESS_SUCCESS: 训练数据预处理完成，等待训练资源中 * TRAINING: 训练中 * TRAIN_FAILED: 训练失败 * TRAIN_SUCCESS: 训练完成，等待预处理资源 * INFERENCE_DATA_PREPROCESSING: 推理数据预处理中 * INFERENCE_DATA_PREPROCESS_FAILED: 推理数据预处理失败 * WAIT_MAIN_FILE_UPLOAD：等待主文件上传 * JOB_SUCCESS: 完成 * WAIT_USER_CONFIRM：等待用户确认训练效果 * JOB_REJECT：驳回 * JOB_PENDING：挂起 * JOB_FINISH：结束，最终状态，不可再做改变
	State Show2dModelTrainingJobResponseState `json:"state"`

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
	ModelVersion *Show2dModelTrainingJobResponseModelVersion `json:"model_version,omitempty"`

	// 抠图类型。默认是AI。 * AI：AI抠图 * MANUAL：人工抠图
	MattingType *Show2dModelTrainingJobResponseMattingType `json:"matting_type,omitempty"`

	// 分身数字人训练视频下载URL。24小时内有效。
	TrainingVideoDownloadUrl *string `json:"training_video_download_url,omitempty"`

	// 身份证正面照片下载URL。24小时内有效。
	IdCardImage1DownloadUrl *string `json:"id_card_image1_download_url,omitempty"`

	// 身份证反面照片下载URL。24小时内有效。
	IdCardImage2DownloadUrl *string `json:"id_card_image2_download_url,omitempty"`

	// 授权书下载URL。24小时内有效。
	GrantFileDownloadUrl *string `json:"grant_file_download_url,omitempty"`

	// 操作日志列表。
	OperationLogs *[]OperationLogInfo `json:"operation_logs,omitempty"`

	// 评论记录列表。
	CommentLogs *[]CommentLogInfo `json:"comment_logs,omitempty"`

	// 遮罩文件是否已上传。
	IsMaskFileUploaded *bool `json:"is_mask_file_uploaded,omitempty"`

	// 遮罩下载URL。24小时内有效。
	MaskFileDownloadUrl *string `json:"mask_file_download_url,omitempty"`

	// 制作审核视频
	VerifyVideoDownloadUrl *string `json:"verify_video_download_url,omitempty"`

	// 标注视频url下载链接。24小时内有效。
	MarkableVideoDownloadUrl *string `json:"markable_video_download_url,omitempty"`

	InferenceDataProcessVideoMarkInfo *InferenceVideoMarkInfo `json:"inference_data_process_video_mark_info,omitempty"`

	InferenceDataProcessActionMarkInfo *InferenceActionMarkInfo `json:"inference_data_process_action_mark_info,omitempty"`

	// 分身数字人是否需要背景替换。需要背景替换的分身数字人训练视频需要绿幕拍摄。
	IsBackgroundReplacement *bool `json:"is_background_replacement,omitempty"`

	// 分身数字人模型分辨率。默认是1080P。 * 1080P：1080P。支持1080P及720P的视频输出。 * 4K：4K。支持4K、1080P及720P的视频输出。
	ModelResolution *string `json:"model_resolution,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o Show2dModelTrainingJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Show2dModelTrainingJobResponse struct{}"
	}

	return strings.Join([]string{"Show2dModelTrainingJobResponse", string(data)}, " ")
}

type Show2dModelTrainingJobResponseState struct {
	value string
}

type Show2dModelTrainingJobResponseStateEnum struct {
	WAIT_FILE_UPLOAD                 Show2dModelTrainingJobResponseState
	AUTO_VERIFYING                   Show2dModelTrainingJobResponseState
	AUTO_VERIFY_FAILED               Show2dModelTrainingJobResponseState
	MANUAL_VERIFYING                 Show2dModelTrainingJobResponseState
	MANUAL_VERIFY_FAILED             Show2dModelTrainingJobResponseState
	MANUAL_VERIFY_SUCCESS            Show2dModelTrainingJobResponseState
	TRAINING_DATA_PREPROCESSING      Show2dModelTrainingJobResponseState
	TRAINING_DATA_PREPROCESS_FAILED  Show2dModelTrainingJobResponseState
	TRAINING_DATA_PREPROCESS_SUCCESS Show2dModelTrainingJobResponseState
	TRAINING                         Show2dModelTrainingJobResponseState
	TRAIN_FAILED                     Show2dModelTrainingJobResponseState
	TRAIN_SUCCESS                    Show2dModelTrainingJobResponseState
	INFERENCE_DATA_PREPROCESSING     Show2dModelTrainingJobResponseState
	INFERENCE_DATA_PREPROCESS_FAILED Show2dModelTrainingJobResponseState
	WAIT_MASK_UPLOAD                 Show2dModelTrainingJobResponseState
	WAIT_MAIN_FILE_UPLOAD            Show2dModelTrainingJobResponseState
	JOB_SUCCESS                      Show2dModelTrainingJobResponseState
	WAIT_USER_CONFIRM                Show2dModelTrainingJobResponseState
	JOB_REJECT                       Show2dModelTrainingJobResponseState
	JOB_PENDING                      Show2dModelTrainingJobResponseState
	JOB_FINISH                       Show2dModelTrainingJobResponseState
}

func GetShow2dModelTrainingJobResponseStateEnum() Show2dModelTrainingJobResponseStateEnum {
	return Show2dModelTrainingJobResponseStateEnum{
		WAIT_FILE_UPLOAD: Show2dModelTrainingJobResponseState{
			value: "WAIT_FILE_UPLOAD",
		},
		AUTO_VERIFYING: Show2dModelTrainingJobResponseState{
			value: "AUTO_VERIFYING",
		},
		AUTO_VERIFY_FAILED: Show2dModelTrainingJobResponseState{
			value: "AUTO_VERIFY_FAILED",
		},
		MANUAL_VERIFYING: Show2dModelTrainingJobResponseState{
			value: "MANUAL_VERIFYING",
		},
		MANUAL_VERIFY_FAILED: Show2dModelTrainingJobResponseState{
			value: "MANUAL_VERIFY_FAILED",
		},
		MANUAL_VERIFY_SUCCESS: Show2dModelTrainingJobResponseState{
			value: "MANUAL_VERIFY_SUCCESS",
		},
		TRAINING_DATA_PREPROCESSING: Show2dModelTrainingJobResponseState{
			value: "TRAINING_DATA_PREPROCESSING",
		},
		TRAINING_DATA_PREPROCESS_FAILED: Show2dModelTrainingJobResponseState{
			value: "TRAINING_DATA_PREPROCESS_FAILED",
		},
		TRAINING_DATA_PREPROCESS_SUCCESS: Show2dModelTrainingJobResponseState{
			value: "TRAINING_DATA_PREPROCESS_SUCCESS",
		},
		TRAINING: Show2dModelTrainingJobResponseState{
			value: "TRAINING",
		},
		TRAIN_FAILED: Show2dModelTrainingJobResponseState{
			value: "TRAIN_FAILED",
		},
		TRAIN_SUCCESS: Show2dModelTrainingJobResponseState{
			value: "TRAIN_SUCCESS",
		},
		INFERENCE_DATA_PREPROCESSING: Show2dModelTrainingJobResponseState{
			value: "INFERENCE_DATA_PREPROCESSING",
		},
		INFERENCE_DATA_PREPROCESS_FAILED: Show2dModelTrainingJobResponseState{
			value: "INFERENCE_DATA_PREPROCESS_FAILED",
		},
		WAIT_MASK_UPLOAD: Show2dModelTrainingJobResponseState{
			value: "WAIT_MASK_UPLOAD",
		},
		WAIT_MAIN_FILE_UPLOAD: Show2dModelTrainingJobResponseState{
			value: "WAIT_MAIN_FILE_UPLOAD",
		},
		JOB_SUCCESS: Show2dModelTrainingJobResponseState{
			value: "JOB_SUCCESS",
		},
		WAIT_USER_CONFIRM: Show2dModelTrainingJobResponseState{
			value: "WAIT_USER_CONFIRM",
		},
		JOB_REJECT: Show2dModelTrainingJobResponseState{
			value: "JOB_REJECT",
		},
		JOB_PENDING: Show2dModelTrainingJobResponseState{
			value: "JOB_PENDING",
		},
		JOB_FINISH: Show2dModelTrainingJobResponseState{
			value: "JOB_FINISH",
		},
	}
}

func (c Show2dModelTrainingJobResponseState) Value() string {
	return c.value
}

func (c Show2dModelTrainingJobResponseState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *Show2dModelTrainingJobResponseState) UnmarshalJSON(b []byte) error {
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

type Show2dModelTrainingJobResponseModelVersion struct {
	value string
}

type Show2dModelTrainingJobResponseModelVersionEnum struct {
	V2 Show2dModelTrainingJobResponseModelVersion
	V3 Show2dModelTrainingJobResponseModelVersion
}

func GetShow2dModelTrainingJobResponseModelVersionEnum() Show2dModelTrainingJobResponseModelVersionEnum {
	return Show2dModelTrainingJobResponseModelVersionEnum{
		V2: Show2dModelTrainingJobResponseModelVersion{
			value: "V2",
		},
		V3: Show2dModelTrainingJobResponseModelVersion{
			value: "V3",
		},
	}
}

func (c Show2dModelTrainingJobResponseModelVersion) Value() string {
	return c.value
}

func (c Show2dModelTrainingJobResponseModelVersion) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *Show2dModelTrainingJobResponseModelVersion) UnmarshalJSON(b []byte) error {
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

type Show2dModelTrainingJobResponseMattingType struct {
	value string
}

type Show2dModelTrainingJobResponseMattingTypeEnum struct {
	AI     Show2dModelTrainingJobResponseMattingType
	MANUAL Show2dModelTrainingJobResponseMattingType
}

func GetShow2dModelTrainingJobResponseMattingTypeEnum() Show2dModelTrainingJobResponseMattingTypeEnum {
	return Show2dModelTrainingJobResponseMattingTypeEnum{
		AI: Show2dModelTrainingJobResponseMattingType{
			value: "AI",
		},
		MANUAL: Show2dModelTrainingJobResponseMattingType{
			value: "MANUAL",
		},
	}
}

func (c Show2dModelTrainingJobResponseMattingType) Value() string {
	return c.value
}

func (c Show2dModelTrainingJobResponseMattingType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *Show2dModelTrainingJobResponseMattingType) UnmarshalJSON(b []byte) error {
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
