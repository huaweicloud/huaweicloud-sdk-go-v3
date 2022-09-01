package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ReqDataByIdCardImage struct {

	// 身份证人像面图像数据，使用base64编码，要求base64编码后大小不超过4M。图像各边的像素大小在300到4000之间，支持JPG格式。
	IdcardImage1 string `json:"idcard_image1" xml:"idcard_image1"`

	// 身份证国徽面图像数据，使用base64编码，要求base64编码后大小不超过4M。图像各边的像素大小在300到4000之间，支持JPG格式。
	IdcardImage2 *string `json:"idcard_image2,omitempty" xml:"idcard_image2"`

	// 现场人像图像数据，使用base64编码，要求base64编码后大小不超过4M。图像各边的像素大小在300到4000之间，支持JPG格式。
	FaceImage string `json:"face_image" xml:"face_image"`
}

func (o ReqDataByIdCardImage) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReqDataByIdCardImage struct{}"
	}

	return strings.Join([]string{"ReqDataByIdCardImage", string(data)}, " ")
}
