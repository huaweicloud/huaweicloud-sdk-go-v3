package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListNodeEncryptdatasResponse Response Object
type ListNodeEncryptdatasResponse struct {

	// 加密数据列表详情
	EncryptDatas *[]EncryptData `json:"encrypt_datas,omitempty"`

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
