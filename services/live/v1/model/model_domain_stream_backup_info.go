package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DomainStreamBackupInfo struct {

	// **参数解释**： 直播推流域名 **约束限制**： 不涉及 **取值范围**： 字符长度为[1-255]位 **默认取值**： 不涉及
	PublishDomain string `json:"publish_domain"`

	// **参数解释**： 主备流开关 **约束限制**： 不涉及 **取值范围**： - true： 开启主备流功能 - false：关闭主备流功能 **默认取值**： false
	StreamBackupEnable bool `json:"stream_backup_enable"`
}

func (o DomainStreamBackupInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DomainStreamBackupInfo struct{}"
	}

	return strings.Join([]string{"DomainStreamBackupInfo", string(data)}, " ")
}
