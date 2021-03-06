package model

import (
	"encoding/json"

	"strings"
)

//
type ConfigByOption struct {
	// 密码强度策略的正则表达式。(当option为password_regex时返回)

	PasswordRegex *string `json:"password_regex,omitempty"`
	// 密码强度策略的描述。(当option为password_regex_description时返回)

	PasswordRegexDescription *string `json:"password_regex_description,omitempty"`
}

func (o ConfigByOption) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ConfigByOption struct{}"
	}

	return strings.Join([]string{"ConfigByOption", string(data)}, " ")
}
