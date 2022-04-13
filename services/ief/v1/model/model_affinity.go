package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 应用实例亲和性规则，仅铂金版支持
type Affinity struct {
	NodeAffinity *AffinityNodeAffinity `json:"nodeAffinity,omitempty"`

	PodAffinity *AffinityPodAffinity `json:"podAffinity,omitempty"`

	PodAntiAffinity *AffinityPodAntiAffinity `json:"podAntiAffinity,omitempty"`
}

func (o Affinity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Affinity struct{}"
	}

	return strings.Join([]string{"Affinity", string(data)}, " ")
}
