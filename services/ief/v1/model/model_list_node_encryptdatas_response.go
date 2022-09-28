package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListNodeEncryptdatasResponse struct {

	// 加密数据列表详情
	EncryptData *[]EncryptData `json:"encrypt_data,omitempty"`

	// 加密数据总数
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListNodeEncryptdatasResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNodeEncryptdatasResponse struct{}"
	}

	return strings.Join([]string{"ListNodeEncryptdatasResponse", string(data)}, " ")
}
