package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DedicatedStorageInfo struct {

	// 专属资源池存储资源规格码。
	SpecCode string `json:"spec_code"`

	// 专属资源池存储主机数量。
	HostNum int32 `json:"host_num"`
}

func (o DedicatedStorageInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DedicatedStorageInfo struct{}"
	}

	return strings.Join([]string{"DedicatedStorageInfo", string(data)}, " ")
}
