package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCertificateTagResponse Response Object
type CreateCertificateTagResponse struct {

	// 请求结果。
	Desc           *string `json:"desc,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateCertificateTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCertificateTagResponse struct{}"
	}

	return strings.Join([]string{"CreateCertificateTagResponse", string(data)}, " ")
}
