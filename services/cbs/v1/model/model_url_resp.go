package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UrlResp
type UrlResp struct {

	// 图片id
	Id string `json:"id"`

	//
	Name *string `json:"name,omitempty"`

	// 访问地址
	Url string `json:"url"`
}

func (o UrlResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UrlResp struct{}"
	}

	return strings.Join([]string{"UrlResp", string(data)}, " ")
}
