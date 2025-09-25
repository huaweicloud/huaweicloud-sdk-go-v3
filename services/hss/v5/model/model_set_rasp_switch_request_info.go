package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SetRaspSwitchRequestInfo struct {

	// **参数解释**: 服务器ID列表，仅支持Linux服务器，要求服务器已开启网页防篡改防护。 **约束限制**: 需要使用 ListWtpProtectHost 接口查询网页防篡改主机防护状态列表信息，在 ListWtpProtectHost 接口的响应体中，os_type 等于 Linux 的 host_id 是符合修改条件的服务器ID。 **取值范围**: 最少1条，最多200条 **默认取值**: 不涉及
	HostIdList []string `json:"host_id_list"`

	// **参数解释**: 动态网页防篡改开启状态，仅支持Linux服务器。 **约束限制**: 不涉及 **取值范围**: - True ：开启动态网页防篡改。 - False ：关闭动态网页防篡改。  **默认取值**: False
	Status bool `json:"status"`
}

func (o SetRaspSwitchRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetRaspSwitchRequestInfo struct{}"
	}

	return strings.Join([]string{"SetRaspSwitchRequestInfo", string(data)}, " ")
}
