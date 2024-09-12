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

	// 任务的状态。 * WAIT_FILE_UPLOAD: 待上传文件 * AUTO_VERIFYING: 自动审核中 * AUTO_VERIFY_FAILED: 自动审核失败 * MANUAL_VERIFYING: 人工审核中 * WAIT_TRAINING_DATA_PREPROCESS: 人工审核中 * MANUAL_VERIFY_FAILED: 人工审核失败 * MANUAL_VERIFY_SUCCESS: 审核通过，等待预处理资源 * TRAINING_DATA_PREPROCESSING: 训练数据预处理中 * TRAINING_DATA_PREPROCESS_FAILED: 训练数据预处理失败 * TRAINING_DATA_PREPROCESS_SUCCESS: 训练数据预处理完成，等待训练资源中 * TRAINING: 训练中 * TRAIN_FAILED: 训练失败 * TRAIN_SUCCESS: 训练完成，等待预处理资源 * INFERENCE_DATA_PREPROCESSING: 推理数据预处理中 * INFERENCE_DATA_PREPROCESS_FAILED: 推理数据预处理失败 * WAIT_MASK_UPLOAD: 等待遮罩上传 * WAIT_MAIN_FILE_UPLOAD: 等待主文件上传 * JOB_SUCCESS: 训练任务完成 * MANUAL_STOP_INFERENCE_DATA_PREPROCESS: 人工中止推理预处理 * MANUAL_STOP_TRAIN: 人工中止训练 * MANUAL_STOP_TRAINING_DATA_PREPROCESS: 人工中止训练预处理 * WAIT_USER_CONFIRM: 等待用户确认训练效果 * JOB_REJECT: 驳回任务 * JOB_PENDING: 挂起任务 * WAIT_ADMIN_CONFIRM: 等待管理员审核 * JOB_FINISH: 任务结束，是最终状态，不支持修改此状态。 * COMPILING: 转编译中 * WAIT_COMPILE: 等待转编译 * COMPILE_FAILED: 转编译失败 * WAIT_GENERATE_ACTION: 等待原子动作生成 * WAIT_ARRANGE: 等待编排 * ACTION_GENERATE_DATA_PROCESSING: 原子动作生成中 * MANUAL_STOP_ACTION_GENERATE_DATA_PROCESSING: 人工中止动作生成 * MANUAL_STOP_ACTION_GENERATE_ORI_PROCESSING: 人工中止动作编排 * ACTION_GENERATE_ORI_PROCESSING: 动作编排中 * ACTION_GENERATE_DATA_FAILED: 动作生成失败 * ACTION_GENERATE_ORI_FAILED: 动作编排失败 * ACTION_GENERATE_ORI_SUCCESS: 动作编排成功 * GENERATE_ACTION_PREPROCESS_FAILED: 动作编排失败 * WAIT_ADMIN_CALIBRATION: 等待管理员确认动作信息
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

	// 分身数字人模型版本。默认是V3版本模型。 * V2: V2版本模型 * V3：V3版本模型 * V3.2：V3.2版本模型
	ModelVersion *Show2dModelTrainingJobResponseModelVersion `json:"model_version,omitempty"`

	// 抠图类型。默认是AI。 * AI：AI抠图 * MANUAL：人工抠图
	MattingType *Show2dModelTrainingJobResponseMattingType `json:"matting_type,omitempty"`

	// 分身数字人模型分辨率。默认是1080P。 * 1080P：1080P。支持1080P及720P的视频输出。 * 4K：4K。支持4K、1080P及720P的视频输出。
	ModelResolution *string `json:"model_resolution,omitempty"`

	// 自定义用户id（如创建任务时设置了X-App-UserId则会携带）。
	AppUserId *string `json:"app_user_id,omitempty"`

	// 是否是基础版的形象训练
	IsFlexus *bool `json:"is_flexus,omitempty"`

	// 分身数字人训练视频下载URL。24小时内有效。
	TrainingVideoDownloadUrl *string `json:"training_video_download_url,omitempty"`

	// 身份证正面照片下载URL。24小时内有效。
	IdCardImage1DownloadUrl *string `json:"id_card_image1_download_url,omitempty"`

	// 身份证反面照片下载URL。24小时内有效。
	IdCardImage2DownloadUrl *string `json:"id_card_image2_download_url,omitempty"`

	// 授权书下载URL。24小时内有效。
	GrantFileDownloadUrl *string `json:"grant_file_download_url,omitempty"`

	// 动作视频
	ActionVideoDownloadUrl *string `json:"action_video_download_url,omitempty"`

	// 音频文件下载url。
	AudioFileDownloadUrl *string `json:"audio_file_download_url,omitempty"`

	// 操作日志列表。
	OperationLogs *[]OperationLogInfo `json:"operation_logs,omitempty"`

	// 生成抠图验证视频时不抠图区域。
	VerifyVideoMattingInfo *[]VerifyVideoMattingInfo `json:"verify_video_matting_info,omitempty"`

	// 评论记录列表。
	CommentLogs *[]CommentLogInfo `json:"comment_logs,omitempty"`

	// 动作视频样例。
	Samples *[]ActionSampleInfo `json:"samples,omitempty"`

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

	InferenceDataProcessEyeCorrectionMarkInfo *InferenceEyeCorrectionMarkInfo `json:"inference_data_process_eye_correction_mark_info,omitempty"`

	// 分身数字人是否需要背景替换。需要背景替换的分身数字人训练视频需要绿幕拍摄。
	IsBackgroundReplacement *bool `json:"is_background_replacement,omitempty"`

	// 转编译任务机型
	WorkerType *[]string `json:"worker_type,omitempty"`

	// 声音训练任务id。
	VoiceTrainJobId *string `json:"voice_train_job_id,omitempty"`

	// flexus版本任务剩余可以重训的次数，每重训一次减1，减到0时不可再重训。
	FlexusRetryCount *int32 `json:"flexus_retry_count,omitempty"`

	// 声音来源类型 * VIDEO：视频中抽取音频 * AUDIO：单独上传的音频
	AudioSourceType *Show2dModelTrainingJobResponseAudioSourceType `json:"audio_source_type,omitempty"`

	// 该任务所生成的模型支持的业务类型，可多选
	SupportedService *[]SupportedServiceEnum `json:"supported_service,omitempty"`

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
	WAIT_FILE_UPLOAD                            Show2dModelTrainingJobResponseState
	AUTO_VERIFYING                              Show2dModelTrainingJobResponseState
	AUTO_VERIFY_FAILED                          Show2dModelTrainingJobResponseState
	MANUAL_VERIFYING                            Show2dModelTrainingJobResponseState
	WAIT_TRAINING_DATA_PREPROCESS               Show2dModelTrainingJobResponseState
	MANUAL_VERIFY_FAILED                        Show2dModelTrainingJobResponseState
	MANUAL_VERIFY_SUCCESS                       Show2dModelTrainingJobResponseState
	TRAINING_DATA_PREPROCESSING                 Show2dModelTrainingJobResponseState
	TRAINING_DATA_PREPROCESS_FAILED             Show2dModelTrainingJobResponseState
	TRAINING_DATA_PREPROCESS_SUCCESS            Show2dModelTrainingJobResponseState
	TRAINING                                    Show2dModelTrainingJobResponseState
	TRAIN_FAILED                                Show2dModelTrainingJobResponseState
	TRAIN_SUCCESS                               Show2dModelTrainingJobResponseState
	INFERENCE_DATA_PREPROCESSING                Show2dModelTrainingJobResponseState
	INFERENCE_DATA_PREPROCESS_FAILED            Show2dModelTrainingJobResponseState
	WAIT_MASK_UPLOAD                            Show2dModelTrainingJobResponseState
	WAIT_MAIN_FILE_UPLOAD                       Show2dModelTrainingJobResponseState
	JOB_SUCCESS                                 Show2dModelTrainingJobResponseState
	MANUAL_STOP_INFERENCE_DATA_PREPROCESS       Show2dModelTrainingJobResponseState
	MANUAL_STOP_TRAIN                           Show2dModelTrainingJobResponseState
	MANUAL_STOP_TRAINING_DATA_PREPROCESS        Show2dModelTrainingJobResponseState
	MANUAL_STOP_BEAUTY_PREPROCESS               Show2dModelTrainingJobResponseState
	WAIT_USER_CONFIRM                           Show2dModelTrainingJobResponseState
	JOB_REJECT                                  Show2dModelTrainingJobResponseState
	JOB_PENDING                                 Show2dModelTrainingJobResponseState
	WAIT_ADMIN_CONFIRM                          Show2dModelTrainingJobResponseState
	JOB_FINISH                                  Show2dModelTrainingJobResponseState
	COMPILING                                   Show2dModelTrainingJobResponseState
	WAIT_COMPILE                                Show2dModelTrainingJobResponseState
	COMPILE_FAILED                              Show2dModelTrainingJobResponseState
	WAIT_BEAUTY                                 Show2dModelTrainingJobResponseState
	WAIT_GENERATE_ACTION                        Show2dModelTrainingJobResponseState
	WAIT_ARRANGE                                Show2dModelTrainingJobResponseState
	ACTION_GENERATE_DATA_PROCESSING             Show2dModelTrainingJobResponseState
	MANUAL_STOP_ACTION_GENERATE_DATA_PROCESSING Show2dModelTrainingJobResponseState
	MANUAL_STOP_ACTION_GENERATE_ORI_PROCESSING  Show2dModelTrainingJobResponseState
	ACTION_GENERATE_ORI_PROCESSING              Show2dModelTrainingJobResponseState
	ACTION_GENERATE_DATA_FAILED                 Show2dModelTrainingJobResponseState
	ACTION_GENERATE_ORI_FAILED                  Show2dModelTrainingJobResponseState
	ACTION_GENERATE_ORI_SUCCESS                 Show2dModelTrainingJobResponseState
	GENERATE_ACTION_PREPROCESS_FAILED           Show2dModelTrainingJobResponseState
	WAIT_ADMIN_CALIBRATION                      Show2dModelTrainingJobResponseState
	BEAUTY_VIDEO_FILE_UPLOADED                  Show2dModelTrainingJobResponseState
	BEAUTYFACE_SUCCESS                          Show2dModelTrainingJobResponseState
	BEAUTYFACE_FAILED                           Show2dModelTrainingJobResponseState
	WAIT_BEAUTY_VIDEO_FILE_UPLOAD               Show2dModelTrainingJobResponseState
	BEAUTYFACE_ROCESSING                        Show2dModelTrainingJobResponseState
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
		WAIT_TRAINING_DATA_PREPROCESS: Show2dModelTrainingJobResponseState{
			value: "WAIT_TRAINING_DATA_PREPROCESS",
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
		MANUAL_STOP_INFERENCE_DATA_PREPROCESS: Show2dModelTrainingJobResponseState{
			value: "MANUAL_STOP_INFERENCE_DATA_PREPROCESS",
		},
		MANUAL_STOP_TRAIN: Show2dModelTrainingJobResponseState{
			value: "MANUAL_STOP_TRAIN",
		},
		MANUAL_STOP_TRAINING_DATA_PREPROCESS: Show2dModelTrainingJobResponseState{
			value: "MANUAL_STOP_TRAINING_DATA_PREPROCESS",
		},
		MANUAL_STOP_BEAUTY_PREPROCESS: Show2dModelTrainingJobResponseState{
			value: "MANUAL_STOP_BEAUTY_PREPROCESS",
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
		WAIT_ADMIN_CONFIRM: Show2dModelTrainingJobResponseState{
			value: "WAIT_ADMIN_CONFIRM",
		},
		JOB_FINISH: Show2dModelTrainingJobResponseState{
			value: "JOB_FINISH",
		},
		COMPILING: Show2dModelTrainingJobResponseState{
			value: "COMPILING",
		},
		WAIT_COMPILE: Show2dModelTrainingJobResponseState{
			value: "WAIT_COMPILE",
		},
		COMPILE_FAILED: Show2dModelTrainingJobResponseState{
			value: "COMPILE_FAILED",
		},
		WAIT_BEAUTY: Show2dModelTrainingJobResponseState{
			value: "WAIT_BEAUTY",
		},
		WAIT_GENERATE_ACTION: Show2dModelTrainingJobResponseState{
			value: "WAIT_GENERATE_ACTION",
		},
		WAIT_ARRANGE: Show2dModelTrainingJobResponseState{
			value: "WAIT_ARRANGE",
		},
		ACTION_GENERATE_DATA_PROCESSING: Show2dModelTrainingJobResponseState{
			value: "ACTION_GENERATE_DATA_PROCESSING",
		},
		MANUAL_STOP_ACTION_GENERATE_DATA_PROCESSING: Show2dModelTrainingJobResponseState{
			value: "MANUAL_STOP_ACTION_GENERATE_DATA_PROCESSING",
		},
		MANUAL_STOP_ACTION_GENERATE_ORI_PROCESSING: Show2dModelTrainingJobResponseState{
			value: "MANUAL_STOP_ACTION_GENERATE_ORI_PROCESSING",
		},
		ACTION_GENERATE_ORI_PROCESSING: Show2dModelTrainingJobResponseState{
			value: "ACTION_GENERATE_ORI_PROCESSING",
		},
		ACTION_GENERATE_DATA_FAILED: Show2dModelTrainingJobResponseState{
			value: "ACTION_GENERATE_DATA_FAILED",
		},
		ACTION_GENERATE_ORI_FAILED: Show2dModelTrainingJobResponseState{
			value: "ACTION_GENERATE_ORI_FAILED",
		},
		ACTION_GENERATE_ORI_SUCCESS: Show2dModelTrainingJobResponseState{
			value: "ACTION_GENERATE_ORI_SUCCESS",
		},
		GENERATE_ACTION_PREPROCESS_FAILED: Show2dModelTrainingJobResponseState{
			value: "GENERATE_ACTION_PREPROCESS_FAILED",
		},
		WAIT_ADMIN_CALIBRATION: Show2dModelTrainingJobResponseState{
			value: "WAIT_ADMIN_CALIBRATION",
		},
		BEAUTY_VIDEO_FILE_UPLOADED: Show2dModelTrainingJobResponseState{
			value: "BEAUTY_VIDEO_FILE_UPLOADED",
		},
		BEAUTYFACE_SUCCESS: Show2dModelTrainingJobResponseState{
			value: "BEAUTYFACE_SUCCESS",
		},
		BEAUTYFACE_FAILED: Show2dModelTrainingJobResponseState{
			value: "BEAUTYFACE_FAILED",
		},
		WAIT_BEAUTY_VIDEO_FILE_UPLOAD: Show2dModelTrainingJobResponseState{
			value: "WAIT_BEAUTY_VIDEO_FILE_UPLOAD",
		},
		BEAUTYFACE_ROCESSING: Show2dModelTrainingJobResponseState{
			value: "BEAUTYFACE_ROCESSING",
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
	V2   Show2dModelTrainingJobResponseModelVersion
	V3   Show2dModelTrainingJobResponseModelVersion
	V3_2 Show2dModelTrainingJobResponseModelVersion
}

func GetShow2dModelTrainingJobResponseModelVersionEnum() Show2dModelTrainingJobResponseModelVersionEnum {
	return Show2dModelTrainingJobResponseModelVersionEnum{
		V2: Show2dModelTrainingJobResponseModelVersion{
			value: "V2",
		},
		V3: Show2dModelTrainingJobResponseModelVersion{
			value: "V3",
		},
		V3_2: Show2dModelTrainingJobResponseModelVersion{
			value: "V3.2",
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

type Show2dModelTrainingJobResponseAudioSourceType struct {
	value string
}

type Show2dModelTrainingJobResponseAudioSourceTypeEnum struct {
	VIDEO Show2dModelTrainingJobResponseAudioSourceType
	AUDIO Show2dModelTrainingJobResponseAudioSourceType
}

func GetShow2dModelTrainingJobResponseAudioSourceTypeEnum() Show2dModelTrainingJobResponseAudioSourceTypeEnum {
	return Show2dModelTrainingJobResponseAudioSourceTypeEnum{
		VIDEO: Show2dModelTrainingJobResponseAudioSourceType{
			value: "VIDEO",
		},
		AUDIO: Show2dModelTrainingJobResponseAudioSourceType{
			value: "AUDIO",
		},
	}
}

func (c Show2dModelTrainingJobResponseAudioSourceType) Value() string {
	return c.value
}

func (c Show2dModelTrainingJobResponseAudioSourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *Show2dModelTrainingJobResponseAudioSourceType) UnmarshalJSON(b []byte) error {
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
