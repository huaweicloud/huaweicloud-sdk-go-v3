package model

import (
	"encoding/json"

	"strings"
)

type DomainMapping struct {
	// 直播播放域名
	PullDomain string `json:"pull_domain"`
	// 直播播放域名关联的推流域名
	PushDomain string `json:"push_domain"`
}

func (o DomainMapping) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DomainMapping struct{}"
	}

	return strings.Join([]string{"DomainMapping", string(data)}, " ")
}
