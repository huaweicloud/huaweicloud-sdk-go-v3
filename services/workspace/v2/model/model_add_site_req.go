package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddSiteReq 初始化站点请求。
type AddSiteReq struct {

	// 站点配置信息。
	SiteConfigs []SiteConfigsRequest `json:"site_configs"`
}

func (o AddSiteReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddSiteReq struct{}"
	}

	return strings.Join([]string{"AddSiteReq", string(data)}, " ")
}
