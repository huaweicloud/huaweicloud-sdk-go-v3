package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAgencyInfoRequest Request Object
type ShowAgencyInfoRequest struct {

	// 请求语言类型。默认en-us。 取值范围： - en-us - zh-cn
	XLanguage *string `json:"X-Language,omitempty"`

	// 委托名称。
	AgencyName string `json:"agency_name"`
}

func (o ShowAgencyInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAgencyInfoRequest struct{}"
	}

	return strings.Join([]string{"ShowAgencyInfoRequest", string(data)}, " ")
}
