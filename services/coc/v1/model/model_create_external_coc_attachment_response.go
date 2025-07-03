package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateExternalCocAttachmentResponse Response Object
type CreateExternalCocAttachmentResponse struct {

	// 业务code，0 代表业务成功，其他数值代表有错误，详情请见错误码。
	Code *string `json:"code,omitempty"`

	// 服务编码。
	ProviderCode *string `json:"provider_code,omitempty"`

	// 错误信息，code为0，此值为空；code不为0，此处为具体的错误描述。
	Msg *string `json:"msg,omitempty"`

	Data           *ExternalAttachment `json:"data,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o CreateExternalCocAttachmentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateExternalCocAttachmentResponse struct{}"
	}

	return strings.Join([]string{"CreateExternalCocAttachmentResponse", string(data)}, " ")
}
