package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SmsCountryResp struct {

	// 国家id
	CountryId *int32 `json:"country_id,omitempty"`

	// 国家(英文)
	CountryNameEn *string `json:"country_name_en,omitempty"`

	// 国家(中文)
	CountryNameZh *string `json:"country_name_zh,omitempty"`
}

func (o SmsCountryResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SmsCountryResp struct{}"
	}

	return strings.Join([]string{"SmsCountryResp", string(data)}, " ")
}
