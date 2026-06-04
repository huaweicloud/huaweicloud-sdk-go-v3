package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HttpsTlsVersion **参数解释：** 配置TLS版本及加密套件，TLS版本需要匹配合适的加密套件使用，详情可参考CDN《用户指南》文档的“TLS版本与加密套件”章节。 **约束限制：** 不涉及
type HttpsTlsVersion struct {

	// **参数解释：** 传输层安全性协议，即TLS版本配置 **约束限制：** - 不可全部关闭 - 需开启连续或单个版本号，例：不可仅开启TLS1.0/1.2而关闭TLS1.1 - 多版本开启时，使用“,”拼接，例：TLSv1.1,TLSv1.2 - 配置TLS版本需要先开启HTTPS国际证书  **取值范围：** 目前支持TLSv1.0/1.1/1.2/1.3四个版本的协议  **默认取值：** 不涉及
	TlsVersion string `json:"tls_version"`

	// **参数解释：** 加密算法套件组 **约束限制：** TLS版本仅选择1.0和（或）1.1时，不支持选择强加密套件。 **取值范围：** - default: 默认加密套件 - general: 一般加密套件 - strict: 强加密套件 - custom: 自定义加密套件  **默认取值：** 不涉及
	CipherSuiteGroup string `json:"cipher_suite_group"`

	// **参数解释：** 加密套件 **约束限制：** - 当选择自定义加密套件时，必传 - 选择多个加密套件是以“,”分隔  **取值范围：** - ECDHE-ECDSA-AES128-GCM-SHA256 - ECDHE-ECDSA-AES256-GCM-SHA384 - ECDHE-RSA-AES128-GCM-SHA256 - ECDHE-RSA-AES256-GCM-SHA384 - DHE-RSA-AES128-GCM-SHA256 - DHE-RSA-AES256-GCM-SHA384 - ECDHE-ECDSA-AES128-SHA256 - ECDHE-ECDSA-AES256-SHA384 - ECDHE-RSA-AES128-SHA - ECDHE-RSA-AES128-SHA256 - ECDHE-RSA-AES256-SHA - ECDHE-RSA-AES256-SHA384 - DHE-RSA-AES128-SHA - DHE-RSA-AES128-SHA256 - DHE-RSA-AES256-SHA - DHE-RSA-AES256-SHA256 - AES128-GCM-SHA256 - AES256-GCM-SHA384 - AES128-SHA256 - AES256-SHA256 - ECDHE-ECDSA-CHACHA20-POLY1305 - ECDHE-RSA-CHACHA20-POLY1305 - DHE-RSA-CHACHA20-POLY1305 - ECDHE-ECDSA-AES128-SHA - ECDHE-ECDSA-AES256-SHA - AES128-SHA - AES256-SHA - DES-CBC3-SHA - RC4-SHA  **默认取值：** 不涉及
	CipherSuite *string `json:"cipher_suite,omitempty"`
}

func (o HttpsTlsVersion) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HttpsTlsVersion struct{}"
	}

	return strings.Join([]string{"HttpsTlsVersion", string(data)}, " ")
}
