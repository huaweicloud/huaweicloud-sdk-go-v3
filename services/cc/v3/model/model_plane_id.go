package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PlaneId 中心网络ID。
type PlaneId struct {

	// 资源ID标识符。
	PlaneId string `json:"plane_id"`
}

func (o PlaneId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PlaneId struct{}"
	}

	return strings.Join([]string{"PlaneId", string(data)}, " ")
}
