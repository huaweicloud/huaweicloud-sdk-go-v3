package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ImageHighresolutionMattingInput struct {

	// 任务的输入类型。可选类型为url（指定的文件地址）
	Type string `json:"type"`

	// 任务的输入详情。针对不同的输入类型有不同的配置。高清图像抠图服务目前只支持识别PNG、JPEG、BMP、JPG、TIF格式的图片，只支持单张图片的url输入方式，图像各边的像素大小在1px至10000px之间， URL提供的图片大小不超过20MB
	Data []ImageHighresolutionMattingInputData `json:"data"`
}

func (o ImageHighresolutionMattingInput) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageHighresolutionMattingInput struct{}"
	}

	return strings.Join([]string{"ImageHighresolutionMattingInput", string(data)}, " ")
}
