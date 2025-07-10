package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowShareSpaceConfigResponse Response Object
type ShowShareSpaceConfigResponse struct {

	// 配置名称。
	ConfigName *string `json:"config_name,omitempty"`

	// 配置值，使用Json字符串,'{\"share_space_name\":“协同空间”, \"use_share_password\":是否使用协同密码：true/false, \"allow_anonymous\":是否匿名加入协同:true/false, \"audio_and_video\":是否使用音频,\"AUDIO\"/\"NONE\"/\"AUDIO_AND_VIDEO\",\"keyboard_mouse_ctl\":云桌面是否可用键鼠true/false,\"anonymous_input_ctrl\":是否开启匿名用户键鼠输入权限true/false, \"is_user_confirm_enabled\":是否需要用户应答true/false,\"wait_confirm_time\":等待时间30-180s}。'
	ConfigValue    *string `json:"config_value,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowShareSpaceConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowShareSpaceConfigResponse struct{}"
	}

	return strings.Join([]string{"ShowShareSpaceConfigResponse", string(data)}, " ")
}
