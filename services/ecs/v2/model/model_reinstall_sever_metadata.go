package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type ReinstallSeverMetadata struct {

	// metadata中的表示加密功能的字段，0代表不加密，1代表加密。  该字段不存在时，云硬盘默认为不加密。
	SystemEncrypted *string `json:"__system__encrypted,omitempty"`

	// metadata中的加密cmkid字段，与__system__encrypted配合表示需要加密，cmkid长度固定为36个字节。  > 说明：  - 请参考[查询密钥列表](https://support.huaweicloud.com/api-dew/ListKeys.html)，通过HTTPS请求获取密钥ID。
	SystemCmkid *string `json:"__system__cmkid,omitempty"`

	// 重装云服务器过程中注入用户数据。  支持注入文本、文本文件或gzip文件。注入内容最大长度32KB。注入内容，需要进行base64格式编码。  了解更多用户数据注入请参考[用户数据注入](https://support.huaweicloud.com/usermanual-ecs/zh-cn_topic_0032380449.html)。
	UserData *string `json:"user_data,omitempty"`
}

func (o ReinstallSeverMetadata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReinstallSeverMetadata struct{}"
	}

	return strings.Join([]string{"ReinstallSeverMetadata", string(data)}, " ")
}
