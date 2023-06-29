package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteEncryptdatasRequest Request Object
type DeleteEncryptdatasRequest struct {

	// 加密数据ID
	EncryptdataId string `json:"encryptdata_id"`

	// 铂金版实例ID，专业版实例为空值
	IefInstanceId *string `json:"ief-instance-id,omitempty"`
}

func (o DeleteEncryptdatasRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteEncryptdatasRequest struct{}"
	}

	return strings.Join([]string{"DeleteEncryptdatasRequest", string(data)}, " ")
}
