package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImageUrlResp
type ImageUrlResp struct {

	// 图片id
	Id string `json:"id"`

	//
	Name *string `json:"name,omitempty"`

	// 访问地址
	Url string `json:"url"`

	Resolution *Resolution `json:"resolution"`
}

func (o ImageUrlResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageUrlResp struct{}"
	}

	return strings.Join([]string{"ImageUrlResp", string(data)}, " ")
}
