package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MetadataInstall metadata字段数据结构说明
type MetadataInstall struct {

	// 重装裸金属服务器过程中注入Linux镜像root密码，用户自定义初始化密码。注：修改密码脚本需经Base64编码。建议密码复杂度如下：长度为8-26位。密码至少必须包含大写字母（A-Z）、小写字母（a-z）、数字（0-9）和特殊字符（!@$%^-_=+[{}]:,./?）中的三种
	UserData *string `json:"user_data,omitempty"`

	// metadata中的表示加密功能的字段，0代表不加密，1代表加密。  该字段不存在时，云硬盘默认为不加密。
	SystemEncrypted *string `json:"__system__encrypted,omitempty"`

	// metadata中的加密cmkid字段，与__system__encrypted配合表示需要加密，cmkid长度固定为36个字节。  > 说明：  - 请参考[查询密钥列表](https://support.huaweicloud.com/api-dew/ListKeys.html)，通过HTTPS请求获取密钥ID。
	SystemCmkid *string `json:"__system__cmkid,omitempty"`
}

func (o MetadataInstall) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MetadataInstall struct{}"
	}

	return strings.Join([]string{"MetadataInstall", string(data)}, " ")
}
