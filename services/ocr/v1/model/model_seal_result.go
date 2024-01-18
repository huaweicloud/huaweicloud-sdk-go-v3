package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SealResult
type SealResult struct {

	// 印章信息列表。
	SealList *[]SealList `json:"seal_list,omitempty"`

	// 在输入图片基础上进行印章擦除后的base64编码图片。
	ErasedSealImage *string `json:"erased_seal_image,omitempty"`
}

func (o SealResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SealResult struct{}"
	}

	return strings.Join([]string{"SealResult", string(data)}, " ")
}
