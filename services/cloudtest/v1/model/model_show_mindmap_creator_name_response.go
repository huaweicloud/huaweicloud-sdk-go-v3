package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowMindmapCreatorNameResponse Response Object
type ShowMindmapCreatorNameResponse struct {

	// 接口调用错误码
	Code *string `json:"code,omitempty"`

	// 接口调用返回体
	Data *interface{} `json:"data,omitempty"`

	// 接口调用错误信息
	Message        *string `json:"message,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowMindmapCreatorNameResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMindmapCreatorNameResponse struct{}"
	}

	return strings.Join([]string{"ShowMindmapCreatorNameResponse", string(data)}, " ")
}
