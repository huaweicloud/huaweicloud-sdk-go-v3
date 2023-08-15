package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ReqDataByVideoAndIdCardImage struct {

	// 身份证人像面图像数据，使用base64编码，要求base64编码后大小不超过4M。图像各边的像素大小在300到4000之间，支持JPG格式。
	IdcardImage1 string `json:"idcard_image1"`

	// 身份证国徽面图像数据，使用base64编码，要求base64编码后大小不超过4M。图像各边的像素大小在300到4000之间，支持JPG格式。
	IdcardImage2 *string `json:"idcard_image2,omitempty"`

	// 现场拍摄人像视频数据，使用base64编码，要求base64编码后大小不超过10M。
	Video string `json:"video"`

	// 动作代码顺序列表，英文逗号（,）分隔。建议单动作，目前支持的动作有： • 1：左摇头 • 2：右摇头 • 3：点头 • 4：嘴部动作
	Actions string `json:"actions"`

	// 该参数为点头动作幅度的判断门限，取值范围：[1,90]，默认为10，单位为度。该值设置越大，则越难判断为点头。
	NodThreshold *float64 `json:"nod_threshold,omitempty"`
}

func (o ReqDataByVideoAndIdCardImage) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReqDataByVideoAndIdCardImage struct{}"
	}

	return strings.Join([]string{"ReqDataByVideoAndIdCardImage", string(data)}, " ")
}
