package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UploadCertsRequestBody struct {

	// 证书文件存放的OBS桶（桶类型必须为标准存储或者低频存储，不支持归档存储）。
	BucketName string `json:"bucket_name"`

	// 证书文件对象。证书文件大小不能超过1M。证书名称在4位到32位之间，必须以字母开头，以（.cer|.crt|.rsa|.jks|.pem|.p10|.pfx|.p12|.csr|.der|.keystore）结尾，可以包含字母、数字、中划线、下划线或者小数点，不能包含其他的特殊字符。
	CertsObject string `json:"certs_object"`
}

func (o UploadCertsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadCertsRequestBody struct{}"
	}

	return strings.Join([]string{"UploadCertsRequestBody", string(data)}, " ")
}
