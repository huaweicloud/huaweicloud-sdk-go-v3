package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PluginResources 插件占用的资源量。
type PluginResources struct {
	InvolvedObject *ObjectReference `json:"involvedObject,omitempty"`

	// **参数解释**： 资源对象的副本数。 **取值范围**： 不涉及。
	Replicas *int32 `json:"replicas,omitempty"`

	// **参数解释**： 申请的资源限制。
	Limits map[string]string `json:"limits,omitempty"`

	// **参数解释**： 申请的资源需求。
	Requests map[string]string `json:"requests,omitempty"`
}

func (o PluginResources) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PluginResources struct{}"
	}

	return strings.Join([]string{"PluginResources", string(data)}, " ")
}
