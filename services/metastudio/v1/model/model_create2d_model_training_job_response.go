package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Create2dModelTrainingJobResponse Response Object
type Create2dModelTrainingJobResponse struct {

	// 任务ID。
	JobId *string `json:"job_id,omitempty"`

	// 分身数字人训练视频上传URL。该url在文件上传成功后失效，只能上传一次。注意：视频必须是1080p或者4K分辨率（横、竖屏皆可）的mp4格式，视频长度须大于等于3分钟且小于等于10分钟，否则审核会不通过。 > 通过该地址上传时，需设置content-type为application/octet-stream
	TrainingVideoUploadUrl *[]string `json:"training_video_upload_url,omitempty"`

	// 分身数字人训练视频上传URL。该url在文件上传成功后失效，只能上传一次。注意：视频必须是1080p或者4K分辨率（横、竖屏皆可）的mp4格式，视频长度须大于等于3分钟且小于等于10分钟，否则审核会不通过。 > 通过该地址上传时，需设置content-type为application/octet-stream
	ActionVideoUploadUrl *[]string `json:"action_video_upload_url,omitempty"`

	// 音频数据训练上传URL。该url在文件上传成功后失效，只能上传一次 > 通过该地址上传时，需设置content-type为application/zip。
	AudioUploadUrl *string `json:"audio_upload_url,omitempty"`

	// 模型封面上传URL。该URL在文件上传成功后失效，只能上传一次。 > 通过该地址上传时，需设置content-type为application/octet-stream
	CoverUploadUrl *string `json:"cover_upload_url,omitempty"`

	// 身份证正面照片上传URL。该URL在文件上传成功后失效，只能上传一次。注意：非NA用户必须上传，否则审核会不通过。 > 通过该地址上传时，需设置content-type为application/octet-stream
	IdCardImage1UploadUrl *string `json:"id_card_image1_upload_url,omitempty"`

	// 身份证反面照片上传URL。该URL在文件上传成功后失效，只能上传一次。注意：非NA用户必须上传，否则审核会不通过。 > 通过该地址上传时，需设置content-type为application/octet-stream
	IdCardImage2UploadUrl *string `json:"id_card_image2_upload_url,omitempty"`

	// 授权书上传URL。该URL在文件上传成功后失效，只能上传一次。注意：非NA用户必须上传，否则审核会不通过。 > 通过该地址上传时，需设置content-type为application/octet-stream
	GrantFileUploadUrl *string `json:"grant_file_upload_url,omitempty"`

	// 美白前图片上传url。 > 通过该地址上传时，需设置content-type为application/octet-stream
	PreBeautyImageUploadUrl *string `json:"pre_beauty_image_upload_url,omitempty"`

	// 美白后图片上传url。 > 通过该地址上传时，需设置content-type为application/octet-stream
	PostBeautyImageUploadUrl *string `json:"post_beauty_image_upload_url,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o Create2dModelTrainingJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Create2dModelTrainingJobResponse struct{}"
	}

	return strings.Join([]string{"Create2dModelTrainingJobResponse", string(data)}, " ")
}
