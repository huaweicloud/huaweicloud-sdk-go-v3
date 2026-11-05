package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ImageUploadResult struct {

	// 图片名称（与请求项对应，保序返回）
	ImageName *string `json:"image_name,omitempty"`

	// 后端生成的图片ID（前端PUT后用此ID加入图库）；失败时为空
	ImageId *string `json:"image_id,omitempty"`

	// OBS PUT预签名上传地址；失败时为空
	UploadUrl *string `json:"upload_url,omitempty"`

	// 预签名过期时间（秒）；失败时为空
	Expire *int64 `json:"expire,omitempty"`

	// 是否成功
	Success *bool `json:"success,omitempty"`

	// 失败原因（如格式不支持/大小超限）
	ErrorMsg *string `json:"error_msg,omitempty"`

	// 失败错误码（与error_msg配套，便于请求方按错误码处理）
	ErrorCode *string `json:"error_code,omitempty"`
}

func (o ImageUploadResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageUploadResult struct{}"
	}

	return strings.Join([]string{"ImageUploadResult", string(data)}, " ")
}
