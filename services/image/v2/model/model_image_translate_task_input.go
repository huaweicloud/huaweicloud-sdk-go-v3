package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ImageTranslateTaskInput struct {

	// 任务的输入类型。可选类型有obs（对象存储服务存储的文件），url（指定的文件地址）
	Type string `json:"type"`

	// 任务的输入详情。针对不同的输入类型有不同的配置。图像翻译服务只支持识别PNG、JPEG、BMP、JPG格式的图片。只支持单张图片输入，,分辨率范围为1px-10000px，且长短边比例不能高于100
	Data []ImageTranslateTaskInputData `json:"data"`
}

func (o ImageTranslateTaskInput) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageTranslateTaskInput struct{}"
	}

	return strings.Join([]string{"ImageTranslateTaskInput", string(data)}, " ")
}
