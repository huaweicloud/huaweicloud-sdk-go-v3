package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteUploadImageResponse Response Object
type ExecuteUploadImageResponse struct {

	// 图片id
	Id string `json:"id"`

	//
	Name *string `json:"name,omitempty"`

	// 访问地址
	Url            string `json:"url"`
	HttpStatusCode int    `json:"-"`
}

func (o ExecuteUploadImageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteUploadImageResponse struct{}"
	}

	return strings.Join([]string{"ExecuteUploadImageResponse", string(data)}, " ")
}
