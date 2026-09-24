package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PageInfo struct {

	// 分页时，作为查询下一页的marker，取值为镜像ID。
	NextMarker string `json:"next_marker"`
}

func (o PageInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PageInfo struct{}"
	}

	return strings.Join([]string{"PageInfo", string(data)}, " ")
}
