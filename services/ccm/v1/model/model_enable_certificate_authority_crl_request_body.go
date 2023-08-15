package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EnableCertificateAuthorityCrlRequestBody struct {

	// 吊销列表文件名称。 > 若用户不指定，系统将默认采用当前证书的父CA ID。
	CrlName *string `json:"crl_name,omitempty"`

	// OBS桶名称。 > - 用户必须已创建委托授权，授予PCA服务对OBS的相关权限，具体参见本文档：**证书吊销处理>查看是否具有委托权限**、**证书吊销处理>创建委托**； > - 指定的OBS桶必须存在，否则将报错。
	ObsBucketName string `json:"obs_bucket_name"`

	// CRL更新周期，单位为\"天\"。
	ValidDays int32 `json:"valid_days"`
}

func (o EnableCertificateAuthorityCrlRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableCertificateAuthorityCrlRequestBody struct{}"
	}

	return strings.Join([]string{"EnableCertificateAuthorityCrlRequestBody", string(data)}, " ")
}
