package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssociateSiteConnectionBandwidthRequestBody 关联分支连接带宽的请求体。
type AssociateSiteConnectionBandwidthRequestBody struct {
	SiteConnection *AssociateSiteConnectionBandwidth `json:"site_connection"`
}

func (o AssociateSiteConnectionBandwidthRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateSiteConnectionBandwidthRequestBody struct{}"
	}

	return strings.Join([]string{"AssociateSiteConnectionBandwidthRequestBody", string(data)}, " ")
}
