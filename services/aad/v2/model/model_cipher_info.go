package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CipherInfo cipher
type CipherInfo struct {

	// 套件名称
	Name string `json:"name"`

	// 对应的加密算法
	Algo string `json:"algo"`

	// 中文描述
	DescCn string `json:"desc_cn"`

	// 英文描述
	DescEn string `json:"desc_en"`
}

func (o CipherInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CipherInfo struct{}"
	}

	return strings.Join([]string{"CipherInfo", string(data)}, " ")
}
