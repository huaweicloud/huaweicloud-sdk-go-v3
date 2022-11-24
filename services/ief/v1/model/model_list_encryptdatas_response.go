package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListEncryptdatasResponse struct {

	// 加密数据列表详情
	EncryptDatas *[]EncryptData `json:"encrypt_datas,omitempty"`

	// 加密数据总数
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListEncryptdatasResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEncryptdatasResponse struct{}"
	}

	return strings.Join([]string{"ListEncryptdatasResponse", string(data)}, " ")
}
