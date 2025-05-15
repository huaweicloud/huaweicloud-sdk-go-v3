package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSiteConnectionBandwidthRequestBody 更改分支连接带宽包的请求体。
type UpdateSiteConnectionBandwidthRequestBody struct {
	SiteConnection *UpdateSiteConnectionBandwidth `json:"site_connection"`
}

func (o UpdateSiteConnectionBandwidthRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSiteConnectionBandwidthRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateSiteConnectionBandwidthRequestBody", string(data)}, " ")
}
