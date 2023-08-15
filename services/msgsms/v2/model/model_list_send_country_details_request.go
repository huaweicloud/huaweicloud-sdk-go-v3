package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSendCountryDetailsRequest Request Object
type ListSendCountryDetailsRequest struct {

	// 国家(英文)
	CountryNameEn *string `json:"country_name_en,omitempty"`

	// 国家(中文)
	CountryNameZh *string `json:"country_name_zh,omitempty"`
}

func (o ListSendCountryDetailsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSendCountryDetailsRequest struct{}"
	}

	return strings.Join([]string{"ListSendCountryDetailsRequest", string(data)}, " ")
}
