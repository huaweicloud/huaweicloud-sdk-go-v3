package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExtraInfoList
type ExtraInfoList struct {

	// 表示key值，可能是qq, wechat, alipay及bank等。
	Item *string `json:"item,omitempty"`

	// 表示value值，对应qq, wechat, alipay及bank等的账号。
	Value *string `json:"value,omitempty"`

	// 对应item关联的额外信息，为bank时第一个默认为户名，第二个为开户行，为alipay时第一个默认为账号名。
	Note *[]string `json:"note,omitempty"`
}

func (o ExtraInfoList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExtraInfoList struct{}"
	}

	return strings.Join([]string{"ExtraInfoList", string(data)}, " ")
}
