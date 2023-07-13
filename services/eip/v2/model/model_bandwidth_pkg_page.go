package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BandwidthPkgPage struct {

	// - 链接
	Href string `json:"href"`

	// - 翻页标志
	Rel string `json:"rel"`
}

func (o BandwidthPkgPage) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BandwidthPkgPage struct{}"
	}

	return strings.Join([]string{"BandwidthPkgPage", string(data)}, " ")
}
