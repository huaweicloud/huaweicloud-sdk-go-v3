package model

import (
	"encoding/json"

	"strings"
)

// 运营商信息
type Operator struct {
	// 运营商ID。

	Id *string `json:"id,omitempty"`
	// 运营商名称。  取值范围： - chinamobile：中国移动； - chinaunicom：中国联通； - chinatelecom：中国电信。

	Name *string `json:"name,omitempty"`
	// 运营商国际化名称。

	I18nName *string `json:"i18n_name,omitempty"`
	// 运营商的简写。

	Sa *string `json:"sa,omitempty"`
}

func (o Operator) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "Operator struct{}"
	}

	return strings.Join([]string{"Operator", string(data)}, " ")
}
