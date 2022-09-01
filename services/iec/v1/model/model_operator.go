package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 运营商信息
type Operator struct {

	// 运营商ID。
	Id *string `json:"id,omitempty" xml:"id"`

	// 运营商名称。  取值范围： - chinamobile：中国移动； - chinaunicom：中国联通； - chinatelecom：中国电信。
	Name *string `json:"name,omitempty" xml:"name"`

	// 运营商国际化名称。
	I18nName *string `json:"i18n_name,omitempty" xml:"i18n_name"`

	// 运营商的简写。
	Sa *string `json:"sa,omitempty" xml:"sa"`
}

func (o Operator) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Operator struct{}"
	}

	return strings.Join([]string{"Operator", string(data)}, " ")
}
