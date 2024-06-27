package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChInstancesInfoRsponseInstanceParamGroup 参数组。
type ChInstancesInfoRsponseInstanceParamGroup struct {

	// 参数ID。
	Id *string `json:"id,omitempty"`

	// 参数名。
	Name *string `json:"name,omitempty"`
}

func (o ChInstancesInfoRsponseInstanceParamGroup) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChInstancesInfoRsponseInstanceParamGroup struct{}"
	}

	return strings.Join([]string{"ChInstancesInfoRsponseInstanceParamGroup", string(data)}, " ")
}
