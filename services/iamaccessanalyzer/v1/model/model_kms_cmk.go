package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// KmsCmk KMS密钥。
type KmsCmk struct {

	// 用于加密密钥的授权。
	Grants string `json:"grants"`
}

func (o KmsCmk) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KmsCmk struct{}"
	}

	return strings.Join([]string{"KmsCmk", string(data)}, " ")
}
