package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateMetaCert struct {

	// 证书名称。
	Name string `json:"name"`
}

func (o CreateMetaCert) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMetaCert struct{}"
	}

	return strings.Join([]string{"CreateMetaCert", string(data)}, " ")
}
