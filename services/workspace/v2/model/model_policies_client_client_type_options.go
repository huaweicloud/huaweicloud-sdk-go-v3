package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PoliciesClientClientTypeOptions 终端类型登录管控的选项。
type PoliciesClientClientTypeOptions struct {

	// 是否勾选Windows客户端。取值为： false：表示未勾选。 true：表示勾选。
	ScTypeWindow *bool `json:"sc_type_window,omitempty"`

	// 是否勾选macOS客户端。取值为： false：表示未勾选。 true：表示勾选。
	ScTypeMac *bool `json:"sc_type_mac,omitempty"`

	// 是否勾选Android客户端。取值为： false：表示未勾选。 true：表示勾选。
	ScTypeAndroid *bool `json:"sc_type_android,omitempty"`

	// 是否勾选Linux客户端。取值为： false：表示未勾选。 true：表示勾选。
	ScTypeLinux *bool `json:"sc_type_linux,omitempty"`

	// 是否勾选Web客户端。取值为： false：表示未勾选。 true：表示勾选。
	ScTypeH5 *bool `json:"sc_type_h5,omitempty"`

	// 是否勾选ios客户端。取值为： false：表示未勾选。 true：表示勾选。
	ScTypeIos *bool `json:"sc_type_ios,omitempty"`

	// 是否勾选鸿蒙客户端。取值为： false：表示未勾选。 true：表示勾选。
	ScTypeHarmonyOs *bool `json:"sc_type_harmony_os,omitempty"`

	// 是否勾选全部硬件终端。取值为： false：表示未勾选。 true：表示勾选。
	TcTypeAll *bool `json:"tc_type_all,omitempty"`
}

func (o PoliciesClientClientTypeOptions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoliciesClientClientTypeOptions struct{}"
	}

	return strings.Join([]string{"PoliciesClientClientTypeOptions", string(data)}, " ")
}
