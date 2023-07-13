package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecutePostCreateImagesResponse Response Object
type ExecutePostCreateImagesResponse struct {

	// 图片id
	Id string `json:"id"`

	//
	Name *string `json:"name,omitempty"`

	// 访问地址
	Url            string `json:"url"`
	HttpStatusCode int    `json:"-"`
}

func (o ExecutePostCreateImagesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecutePostCreateImagesResponse struct{}"
	}

	return strings.Join([]string{"ExecutePostCreateImagesResponse", string(data)}, " ")
}
