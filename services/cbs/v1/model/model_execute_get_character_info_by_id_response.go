package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ExecuteGetCharacterInfoByIdResponse struct {

	// 创建时间
	CreateTime string `json:"create_time"`

	// 更新时间
	UpdateTime string `json:"update_time"`

	// 形象的个人姓名
	CharacterName string `json:"character_name"`

	Gender *int32 `json:"gender,omitempty"`

	// 形象id
	Id *string `json:"id,omitempty"`

	// 形象名
	Name *string `json:"name,omitempty"`

	// 形象obs地址
	PhotoUrl *string `json:"photo_url,omitempty"`

	// 姿态： 0：站姿全身 1：站姿半身 2：坐姿全身 3：坐姿半身
	Posture *int32 `json:"posture,omitempty"`

	// 估算的训练结束时间
	TrainFinishTimeEstimate *string `json:"train_finish_time_estimate,omitempty"`

	// 训练开始时间
	TrainStartTime *string `json:"train_start_time,omitempty"`

	// 训练状态： 0：预处理 1：训练中 2：训练成功 3：训练失败 4：预览视频生成中
	TrainStatus *int32 `json:"train_status,omitempty"`

	// 形象类型： 0：预制形象 1：用户自定义形象
	Type *int32 `json:"type,omitempty"`

	// 形象在中心时的图片obs 地址
	CenterPhotoUrl string `json:"center_photo_url"`

	// 合成错误信息
	ErrorMsg *string `json:"error_msg,omitempty"`

	// 形象在左时的图片obs 地址 考虑兼容性：如果为null，形象无法使用左右配置
	LeftPhotoUrl *string `json:"left_photo_url,omitempty"`

	// 预览视频
	PreviewVideoUrl *string `json:"preview_video_url,omitempty"`

	// 形象在右时的图片obs 地址 考虑兼容性：如果为null，形象无法使用左右配置
	RightPhotoUrl *string `json:"right_photo_url,omitempty"`

	// 显示效果最佳预览
	BestImgQualityPreviewUrl *string `json:"best_img_quality_preview_url,omitempty"`

	// 音唇同步最佳预览
	BestLipSyncPreviewUrl *string `json:"best_lip_sync_preview_url,omitempty"`

	// 嘴巴部分效果最佳
	BestMouthRecPreviewUrl *string `json:"best_mouth_rec_preview_url,omitempty"`

	// 是否有人像分割数据
	HaveSegmentData bool `json:"have_segment_data"`

	// 合成原始视频地址
	InitialVideoUrl *string `json:"initial_video_url,omitempty"`

	// 抠图背景地址
	BackgroundUrl *string `json:"background_url,omitempty"`

	// 0: best img quality 1: best lip sync 2: best mouth rec
	Model *int32 `json:"model,omitempty"`

	CharaterPosition *CharacterPosition `json:"charater_position,omitempty"`

	CharaterDimension *CharacterDimension `json:"charater_dimension,omitempty"`
	HttpStatusCode    int                 `json:"-"`
}

func (o ExecuteGetCharacterInfoByIdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteGetCharacterInfoByIdResponse struct{}"
	}

	return strings.Join([]string{"ExecuteGetCharacterInfoByIdResponse", string(data)}, " ")
}
