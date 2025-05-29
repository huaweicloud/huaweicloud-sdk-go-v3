package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateCertificateTagRequestBody struct {
	Tag *ScsResourceTag `json:"tag"`
}

func (o CreateCertificateTagRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCertificateTagRequestBody struct{}"
	}

	return strings.Join([]string{"CreateCertificateTagRequestBody", string(data)}, " ")
}
