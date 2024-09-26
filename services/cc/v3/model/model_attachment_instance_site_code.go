package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AttachmentInstanceSiteCode 中心网络附件对端实例的站点编码。
type AttachmentInstanceSiteCode struct {

	// 中心网络附件对端实例的站点编码。
	AttachmentInstanceSiteCode string `json:"attachment_instance_site_code"`
}

func (o AttachmentInstanceSiteCode) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachmentInstanceSiteCode struct{}"
	}

	return strings.Join([]string{"AttachmentInstanceSiteCode", string(data)}, " ")
}
