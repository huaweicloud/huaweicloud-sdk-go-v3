package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UploadFileRequest Request Object
type UploadFileRequest struct {

	// **参数解释**： 对话路由ID。 **约束限制**： 不涉及 **取值范围**： 仅支持字母、数字、中划线和下划线，长度为[1-64]个字符。 **默认取值**： 不涉及
	XChatRouteId string `json:"X-Chat-Route-Id"`

	Body *UploadFileRequestBody `json:"body,omitempty" type:"multipart"`
}

func (o UploadFileRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadFileRequest struct{}"
	}

	return strings.Join([]string{"UploadFileRequest", string(data)}, " ")
}
