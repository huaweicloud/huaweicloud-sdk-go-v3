package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FsDirUasge 目录资源使用情况(包含子目录)
type FsDirUasge struct {

	// 占用容量，单位：byte
	UsedCapacity int64 `json:"used_capacity"`
}

func (o FsDirUasge) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FsDirUasge struct{}"
	}

	return strings.Join([]string{"FsDirUasge", string(data)}, " ")
}
