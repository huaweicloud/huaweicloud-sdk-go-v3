package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OBS桶信息
type ObsInfo struct {

	// OBS桶名称
	Name *string `json:"name,omitempty" xml:"name"`

	// OBS桶地址
	Addr *string `json:"addr,omitempty" xml:"addr"`
}

func (o ObsInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ObsInfo struct{}"
	}

	return strings.Join([]string{"ObsInfo", string(data)}, " ")
}
