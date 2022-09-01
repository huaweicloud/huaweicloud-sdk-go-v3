package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type V2UserData struct {

	// 采样点观众数。
	Value *int64 `json:"value,omitempty" xml:"value"`

	// 采样时间。日期格式按照ISO8601表示法，并使用UTC时间。 格式为：YYYY-MM-DDThh:mm:ssZ。
	Time *string `json:"time,omitempty" xml:"time"`
}

func (o V2UserData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "V2UserData struct{}"
	}

	return strings.Join([]string{"V2UserData", string(data)}, " ")
}
