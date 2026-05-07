package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDomainStreamBackupRequest Request Object
type ShowDomainStreamBackupRequest struct {

	// **参数解释**： 直播推流域名 **约束限制**： 不涉及 **取值范围**： 字符长度为[1-255]位 **默认取值**： 不涉及
	PublishDomain string `json:"publish_domain"`
}

func (o ShowDomainStreamBackupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDomainStreamBackupRequest struct{}"
	}

	return strings.Join([]string{"ShowDomainStreamBackupRequest", string(data)}, " ")
}
