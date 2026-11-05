package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RemoveImagesResponse Response Object
type RemoveImagesResponse struct {

	// 每张图片的移除结果
	Results        *[]ImageRemoveResult `json:"results,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o RemoveImagesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RemoveImagesResponse struct{}"
	}

	return strings.Join([]string{"RemoveImagesResponse", string(data)}, " ")
}
