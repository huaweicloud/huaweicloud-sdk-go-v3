package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowResSceneResponse struct {

	// 是否成功。
	IsSuccess *bool `json:"is_success,omitempty" xml:"is_success"`

	Scene *ResScene `json:"scene,omitempty" xml:"scene"`

	// 返回消息（请求成功时，不返回此字段）。
	Message *string `json:"message,omitempty" xml:"message"`

	// 错误码（请求成功时，不返回此字段）。
	ErrorCode      *string `json:"error_code,omitempty" xml:"error_code"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowResSceneResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowResSceneResponse struct{}"
	}

	return strings.Join([]string{"ShowResSceneResponse", string(data)}, " ")
}
