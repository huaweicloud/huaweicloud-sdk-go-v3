package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AccessPointModel struct {

	// 局点名字
	Region string `json:"region"`
}

func (o AccessPointModel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccessPointModel struct{}"
	}

	return strings.Join([]string{"AccessPointModel", string(data)}, " ")
}
