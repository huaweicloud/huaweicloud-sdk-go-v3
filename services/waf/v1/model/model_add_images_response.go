package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddImagesResponse Response Object
type AddImagesResponse struct {

	// 每张图片的加入结果
	Results        *[]ImageAddResult `json:"results,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o AddImagesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddImagesResponse struct{}"
	}

	return strings.Join([]string{"AddImagesResponse", string(data)}, " ")
}
