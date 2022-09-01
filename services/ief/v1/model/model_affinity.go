package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 应用实例亲和性规则，仅铂金版支持
type Affinity struct {
	NodeAffinity *AffinityNodeAffinity `json:"nodeAffinity,omitempty" xml:"nodeAffinity"`

	PodAffinity *AffinityPodAffinity `json:"podAffinity,omitempty" xml:"podAffinity"`

	PodAntiAffinity *AffinityPodAntiAffinity `json:"podAntiAffinity,omitempty" xml:"podAntiAffinity"`
}

func (o Affinity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Affinity struct{}"
	}

	return strings.Join([]string{"Affinity", string(data)}, " ")
}
