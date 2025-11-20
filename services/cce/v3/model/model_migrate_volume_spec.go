package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MigrateVolumeSpec 服务器重装云硬盘配置
type MigrateVolumeSpec struct {

	// **参数解释**: 用户主密钥ID。 [> 获取密钥ID的方法请参考：[查询密钥列表](https://support.huaweicloud.com/api-dew/ListKeys.html)](tag:hws) [> 获取密钥ID的方法请参考：[查询密钥列表](https://support.huaweicloud.com/intl/zh-cn/api-dew/ListKeys.html)](tag:hws_hk)  **约束限制**: 不涉及 **取值范围**: 不涉及 **默认取值**: 默认为空，表示云硬盘不加密。
	CmkID *string `json:"cmkID,omitempty"`
}

func (o MigrateVolumeSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MigrateVolumeSpec struct{}"
	}

	return strings.Join([]string{"MigrateVolumeSpec", string(data)}, " ")
}
