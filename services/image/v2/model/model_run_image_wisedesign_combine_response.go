package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type RunImageWisedesignCombineResponse struct {

	// 图片合成后图像的64位编码
	ResultBase64   *string `json:"result_base64,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RunImageWisedesignCombineResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunImageWisedesignCombineResponse struct{}"
	}

	return strings.Join([]string{"RunImageWisedesignCombineResponse", string(data)}, " ")
}
