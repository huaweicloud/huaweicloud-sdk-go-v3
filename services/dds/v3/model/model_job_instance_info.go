package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type JobInstanceInfo struct {

	// 实例ID。
	Id string `json:"id" xml:"id"`

	// 实例名称。
	Name string `json:"name" xml:"name"`
}

func (o JobInstanceInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobInstanceInfo struct{}"
	}

	return strings.Join([]string{"JobInstanceInfo", string(data)}, " ")
}
