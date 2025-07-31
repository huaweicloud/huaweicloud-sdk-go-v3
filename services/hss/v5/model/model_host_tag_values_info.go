package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HostTagValuesInfo struct {

	// 资源标签key值
	Key *string `json:"key,omitempty"`

	// 资源标签values
	Values *[]string `json:"values,omitempty"`
}

func (o HostTagValuesInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HostTagValuesInfo struct{}"
	}

	return strings.Join([]string{"HostTagValuesInfo", string(data)}, " ")
}
