package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ReqDataByIdCardImage struct {

	// 身份证人像面图像数据，使用base64编码，要求base64编码后大小不超过4M。图像各边的像素大小在300到4000之间，支持JPG格式。
	IdcardImage1 string `json:"idcard_image1"`

	// 身份证国徽面图像数据，使用base64编码，要求base64编码后大小不超过4M。图像各边的像素大小在300到4000之间，支持JPG格式。
	IdcardImage2 *string `json:"idcard_image2,omitempty"`

	// 现场人像图像数据，使用base64编码，要求base64编码后大小不超过4M。图像各边的像素大小在300到4000之间，支持JPG格式。
	FaceImage string `json:"face_image"`

	// 响应参数similarity是否详细显示，默认为false。 - true表示响应中的similarity为0~1000的小数。 - false表示响应中的similarity为0~100的整数。
	Detail *bool `json:"detail,omitempty"`

	// 是否允许对入参face_image进行人脸检测及图片裁剪，默认为true，表示允许。
	Crop *bool `json:"crop,omitempty"`
}

func (o ReqDataByIdCardImage) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReqDataByIdCardImage struct{}"
	}

	return strings.Join([]string{"ReqDataByIdCardImage", string(data)}, " ")
}
