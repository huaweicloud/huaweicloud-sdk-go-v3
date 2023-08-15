package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VerifyCertificateDto 验证CA证书结构体。
type VerifyCertificateDto struct {

	// 验证证书的内容信息。
	VerifyContent string `json:"verify_content"`
}

func (o VerifyCertificateDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VerifyCertificateDto struct{}"
	}

	return strings.Join([]string{"VerifyCertificateDto", string(data)}, " ")
}
