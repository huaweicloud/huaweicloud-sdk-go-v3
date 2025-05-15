package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTtsPhoneticSymbolResponse Response Object
type ShowTtsPhoneticSymbolResponse struct {

	// 音标列表
	Data           *[]string `json:"data,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ShowTtsPhoneticSymbolResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTtsPhoneticSymbolResponse struct{}"
	}

	return strings.Join([]string{"ShowTtsPhoneticSymbolResponse", string(data)}, " ")
}
