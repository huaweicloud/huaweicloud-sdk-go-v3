package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PresignUploadUrlBody struct {

	// 待获取预签名上传地址的图片列表
	Images []ImageUploadItem `json:"images"`
}

func (o PresignUploadUrlBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PresignUploadUrlBody struct{}"
	}

	return strings.Join([]string{"PresignUploadUrlBody", string(data)}, " ")
}
