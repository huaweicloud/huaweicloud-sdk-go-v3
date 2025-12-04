package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Metric struct {

	// AOM实例配置。
	Aom *[]Aom `json:"aom,omitempty"`
}

func (o Metric) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Metric struct{}"
	}

	return strings.Join([]string{"Metric", string(data)}, " ")
}
