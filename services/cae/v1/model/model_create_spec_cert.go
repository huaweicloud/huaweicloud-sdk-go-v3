package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateSpecCert struct {

	// 证书内容。
	Crt string `json:"crt"`

	// 私钥内容。
	Key string `json:"key"`
}

func (o CreateSpecCert) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSpecCert struct{}"
	}

	return strings.Join([]string{"CreateSpecCert", string(data)}, " ")
}
