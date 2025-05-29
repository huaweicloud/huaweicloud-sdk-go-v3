package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DataVolumesMetadata
type DataVolumesMetadata struct {

	// metadata中的表示加密功能的字段，0代表不加密，1代表加密。  该字段不存在时，云硬盘默认为不加密。
	SystemEncrypted *string `json:"__system__encrypted,omitempty"`

	// metadata中的加密cmkid字段，与__system__encrypted配合表示需要加密，cmkid长度固定为36个字节。  > 说明：  - 请参考[查询密钥列表](https://support.huaweicloud.com/api-dew/ListKeys.html)，通过HTTPS请求获取密钥ID。
	SystemCmkid *string `json:"__system__cmkid,omitempty"`
}

func (o DataVolumesMetadata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataVolumesMetadata struct{}"
	}

	return strings.Join([]string{"DataVolumesMetadata", string(data)}, " ")
}
