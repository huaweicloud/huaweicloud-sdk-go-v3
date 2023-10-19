package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeSeversOsMetadataWithoutCloudInitOption
type ChangeSeversOsMetadataWithoutCloudInitOption struct {

	// metadata中的表示加密功能的字段，0代表不加密，1代表加密。  该字段不存在时，云硬盘默认为不加密。
	SystemEncrypted *string `json:"__system__encrypted,omitempty"`

	// metadata中的加密cmkid字段，与__system__encrypted配合表示需要加密，cmkid长度固定为36个字节。  > 说明：  - 请参考[查询密钥列表](https://support.huaweicloud.com/api-dew/ListKeys.html)，通过HTTPS请求获取密钥ID。
	SystemCmkid *string `json:"__system__cmkid,omitempty"`

	// 如果您已拥有操作系统或软件的许可证（一般是指按物理插槽数、物理内核数等进行认证的许可证），您可以通过自带许可（BYOL）的方式将业务完整迁移到云平台，继续使用您的许可证。 - true： 使用自有license - 其他值： 视为非法参数，接口报错
	Byol *string `json:"BYOL,omitempty"`
}

func (o ChangeSeversOsMetadataWithoutCloudInitOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeSeversOsMetadataWithoutCloudInitOption struct{}"
	}

	return strings.Join([]string{"ChangeSeversOsMetadataWithoutCloudInitOption", string(data)}, " ")
}
