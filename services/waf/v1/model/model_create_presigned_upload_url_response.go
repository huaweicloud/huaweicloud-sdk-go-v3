package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePresignedUploadUrlResponse Response Object
type CreatePresignedUploadUrlResponse struct {

	// 每张图片的预签名获取结果（与请求项保序）
	Results        *[]ImageUploadResult `json:"results,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o CreatePresignedUploadUrlResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePresignedUploadUrlResponse struct{}"
	}

	return strings.Join([]string{"CreatePresignedUploadUrlResponse", string(data)}, " ")
}
