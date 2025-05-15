package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSiteConnectionBandwidthSizeRequestBody 更改分支连接带宽大小的请求体。
type UpdateSiteConnectionBandwidthSizeRequestBody struct {
	SiteConnection *UpdateSiteConnectionBandwidthSize `json:"site_connection"`
}

func (o UpdateSiteConnectionBandwidthSizeRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSiteConnectionBandwidthSizeRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateSiteConnectionBandwidthSizeRequestBody", string(data)}, " ")
}
