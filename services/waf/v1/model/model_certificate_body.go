package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CertificateBody struct {

	// 证书id
	Id string `json:"id"`

	// 证书名
	Name string `json:"name"`

	// 证书过期时间戳
	ExpireTime *int64 `json:"expire_time,omitempty"`

	// **参数解释：** 证书过期状态 **约束限制：** 不涉及 **取值范围：**  - 0:未过期  - 1:已过期  - 2:即将过期（证书将在一个月内过期）  **默认取值：** 不涉及
	ExpStatus *int32 `json:"exp_status,omitempty"`

	// 证书上传时间戳
	Timestamp int64 `json:"timestamp"`

	// 证书关联的域名信息
	BindHost *[]BindHost `json:"bind_host,omitempty"`
}

func (o CertificateBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CertificateBody struct{}"
	}

	return strings.Join([]string{"CertificateBody", string(data)}, " ")
}
