package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MetaCert struct {

	// 证书ID。
	Id *string `json:"id,omitempty"`

	// 证书名称。
	Name *string `json:"name,omitempty"`
}

func (o MetaCert) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MetaCert struct{}"
	}

	return strings.Join([]string{"MetaCert", string(data)}, " ")
}
