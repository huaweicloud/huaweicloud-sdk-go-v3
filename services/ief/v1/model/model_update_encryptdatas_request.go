package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateEncryptdatasRequest Request Object
type UpdateEncryptdatasRequest struct {

	// 加密数据ID
	EncryptdataId string `json:"encryptdata_id"`

	// 铂金版实例ID，专业版实例为空值
	IefInstanceId *string `json:"ief-instance-id,omitempty"`

	Body *EncryptDataReq `json:"body,omitempty"`
}

func (o UpdateEncryptdatasRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEncryptdatasRequest struct{}"
	}

	return strings.Join([]string{"UpdateEncryptdatasRequest", string(data)}, " ")
}
