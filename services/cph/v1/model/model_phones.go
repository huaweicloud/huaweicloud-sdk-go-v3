package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Phones struct {

	// 手机列表
	Phones []PhonesPhones `json:"phones" xml:"phones"`
}

func (o Phones) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Phones struct{}"
	}

	return strings.Join([]string{"Phones", string(data)}, " ")
}
