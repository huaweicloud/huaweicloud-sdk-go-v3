package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UnbindingGeipBody struct {

	// 全局弹性公网IP的ID
	GlobalEipId string `json:"global_eip_id"`
}

func (o UnbindingGeipBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UnbindingGeipBody struct{}"
	}

	return strings.Join([]string{"UnbindingGeipBody", string(data)}, " ")
}
