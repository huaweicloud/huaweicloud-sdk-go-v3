package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type StandardReqDataByNameAndId struct {

	// 被验证人的姓名。
	VerificationName string `json:"verification_name"`

	// 被验证人的身份证号码。
	VerificationId string `json:"verification_id"`

	// 现场人像图像数据，使用base64编码，要求base64编码后大小不超过4M。图像各边的像素大小在300到4000之间，支持JPG格式。
	FaceImage string `json:"face_image"`

	// 响应参数similarity是否详细显示，默认为false。 - true表示响应中的similarity为0~1000的小数。 - false表示响应中的similarity为0~100的整数。
	Detail *bool `json:"detail,omitempty"`

	// 是否允许对入参face_image进行人脸检测及图片裁剪，默认为true，表示允许。
	Crop *bool `json:"crop,omitempty"`
}

func (o StandardReqDataByNameAndId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StandardReqDataByNameAndId struct{}"
	}

	return strings.Join([]string{"StandardReqDataByNameAndId", string(data)}, " ")
}
