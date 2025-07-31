package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ProjectConfigInfo 配置信息
type ProjectConfigInfo struct {

	// 配置名称，包含如下    - password_min_len ：密码最小长度    - password_digit_min_num ：密码中最少包含数字数量    - password_upper_letter_min_num ：密码中最少包含大写字母数量    - password_lower_letter_min_num ：密码中最少包含小写字母数量    - password_special_char_min_num ：密码中最少包含特殊字符数量    - weakpwd: 自定义弱口令策略
	ConfigName *string `json:"config_name,omitempty"`

	// 配置内容
	ConfigValue *string `json:"config_value,omitempty"`
}

func (o ProjectConfigInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProjectConfigInfo struct{}"
	}

	return strings.Join([]string{"ProjectConfigInfo", string(data)}, " ")
}
