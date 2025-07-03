package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAttachmentResponse Response Object
type CreateAttachmentResponse struct {

	// 业务code，0 代表业务成功，其他数值代表有错误，详情请见错误码。
	Code *string `json:"code,omitempty"`

	// 服务编码。
	ProviderCode *string `json:"provider_code,omitempty"`

	// 错误信息，code为0，此值为空；code不为0，此处为具体的错误描述。
	Msg *string `json:"msg,omitempty"`

	Data *ExternalAttachment `json:"data,omitempty"`

	// **参数解释：** 附件文件名。 **取值范围：** 不涉及 **默认取值：** 不涉及
	Name *string `json:"name,omitempty"`

	// **参数解释：** 附件唯一id。 **取值范围：** 不涉及 **默认取值：** 不涉及
	Id             *string `json:"id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateAttachmentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAttachmentResponse struct{}"
	}

	return strings.Join([]string{"CreateAttachmentResponse", string(data)}, " ")
}
