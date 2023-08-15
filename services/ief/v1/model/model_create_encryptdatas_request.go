package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateEncryptdatasRequest Request Object
type CreateEncryptdatasRequest struct {

	// 铂金版实例ID，专业版实例为空值
	IefInstanceId *string `json:"ief-instance-id,omitempty"`

	Body *EncryptDataReq `json:"body,omitempty"`
}

func (o CreateEncryptdatasRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEncryptdatasRequest struct{}"
	}

	return strings.Join([]string{"CreateEncryptdatasRequest", string(data)}, " ")
}
