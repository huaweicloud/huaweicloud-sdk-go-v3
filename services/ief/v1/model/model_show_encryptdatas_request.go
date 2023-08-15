package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowEncryptdatasRequest Request Object
type ShowEncryptdatasRequest struct {

	// 加密数据ID
	EncryptdataId string `json:"encryptdata_id"`

	// 铂金版实例ID，专业版实例为空值
	IefInstanceId *string `json:"ief-instance-id,omitempty"`
}

func (o ShowEncryptdatasRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEncryptdatasRequest struct{}"
	}

	return strings.Join([]string{"ShowEncryptdatasRequest", string(data)}, " ")
}
