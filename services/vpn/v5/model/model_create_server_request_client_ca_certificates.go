package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateServerRequestClientCaCertificates struct {

	// 证书名
	Name *string `json:"name,omitempty"`

	// 证书内容
	Content string `json:"content"`
}

func (o CreateServerRequestClientCaCertificates) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateServerRequestClientCaCertificates struct{}"
	}

	return strings.Join([]string{"CreateServerRequestClientCaCertificates", string(data)}, " ")
}
