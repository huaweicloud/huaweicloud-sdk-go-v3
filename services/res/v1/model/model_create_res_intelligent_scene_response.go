package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateResIntelligentSceneResponse struct {

	// 是否成功。
	IsSuccess *bool `json:"is_success,omitempty" xml:"is_success"`

	// 返回消息（请求成功时，不返回此字段）。
	Message *string `json:"message,omitempty" xml:"message"`

	// 错误码（请求成功时，不返回此字段）。
	ErrorCode *string `json:"error_code,omitempty" xml:"error_code"`

	Scene          *Scene `json:"scene,omitempty" xml:"scene"`
	HttpStatusCode int    `json:"-"`
}

func (o CreateResIntelligentSceneResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateResIntelligentSceneResponse struct{}"
	}

	return strings.Join([]string{"CreateResIntelligentSceneResponse", string(data)}, " ")
}
