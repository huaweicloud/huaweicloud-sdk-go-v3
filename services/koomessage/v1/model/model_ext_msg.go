package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ExtMsg struct {

	// 打开方式。 - 0：webView打开  - 1：浏览器打开   > action_type=OPEN_URL必填，其他不填。
	OpenInBrowser *string `json:"open_in_browser,omitempty"`

	// 标题，必填，长度范围为1-20个字符。 > action_type=OPEN_URL必填，其他不填。
	WebTitle *string `json:"web_title,omitempty"`

	// app包名，长度范围为1-50个字符。 > action_type=OPEN_APP必填，其他不填。
	PackageName *string `json:"package_name,omitempty"`

	// 商家应用的appid，长度范围为0-60个字符。 > action_type=OPEN_APP必填，其他不填。
	AppId *string `json:"app_id,omitempty"`

	// 兜底url，长度范围为0-1000个字符，支持http/https。 > action_type=OPEN_APP选填，其他不填。
	BrowserFloorUrl *string `json:"browser_floor_url,omitempty"`

	// 依赖的快应用引擎版本号，长度范围为1-50个字符。 > action_type=OPEN_QUICK必填，其他不填。
	DependEngineVer *string `json:"depend_engine_ver,omitempty"`

	// 第三方服务名，长度范围为1-50个字符。 > action_type=OPEN_QUICK必填，其他不填。
	ThirdServiceName *string `json:"third_service_name,omitempty"`
}

func (o ExtMsg) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExtMsg struct{}"
	}

	return strings.Join([]string{"ExtMsg", string(data)}, " ")
}
