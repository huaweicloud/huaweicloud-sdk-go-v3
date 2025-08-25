package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AppResourceInfo struct {

	// 当前租户在cpcs创建的简单应用数量
	AppNum int32 `json:"app_num"`
}

func (o AppResourceInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AppResourceInfo struct{}"
	}

	return strings.Join([]string{"AppResourceInfo", string(data)}, " ")
}
