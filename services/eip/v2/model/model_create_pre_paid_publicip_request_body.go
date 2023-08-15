package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePrePaidPublicipRequestBody 创建包周期的弹性公网IP
type CreatePrePaidPublicipRequestBody struct {
	Publicip *CreatePrePaidPublicipOption `json:"publicip"`

	Bandwidth *CreatePublicipBandwidthOption `json:"bandwidth"`

	ExtendParam *CreatePrePaidPublicipExtendParamOption `json:"extendParam,omitempty"`

	// 企业项目ID。最大长度36字节，带“-”连字符的UUID格式，或者是字符串“0”。  创建弹性公网IP时，给弹性公网IP绑定企业项目ID。  不指定该参数时，默认值是 0
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o CreatePrePaidPublicipRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePrePaidPublicipRequestBody struct{}"
	}

	return strings.Join([]string{"CreatePrePaidPublicipRequestBody", string(data)}, " ")
}
