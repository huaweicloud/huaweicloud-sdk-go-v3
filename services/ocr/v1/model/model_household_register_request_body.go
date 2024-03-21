package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HouseholdRegisterRequestBody
type HouseholdRegisterRequestBody struct {

	// 与url二选一。 图像数据，base64编码，要求base64编码后大小不超过10MB。图片最小边不小于15px，最长边不超过8192px，支持JPEG、JPG、PNG、BMP、TIFF格式。
	Image *string `json:"image,omitempty"`

	// 与image二选一 图片的URL路径，目前仅支持华为云上OBS提供的匿名公开授权访问的URL以及公网URL。
	Url *string `json:"url,omitempty"`
}

func (o HouseholdRegisterRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HouseholdRegisterRequestBody struct{}"
	}

	return strings.Join([]string{"HouseholdRegisterRequestBody", string(data)}, " ")
}
