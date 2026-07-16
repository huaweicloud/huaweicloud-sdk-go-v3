package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResourceFlavorStatus 资源规格状态信息。
type ResourceFlavorStatus struct {

	// **参数解释**：资源规格的容量状态，格式为key/value键值对。其中，key为az编码，value为对应az资源的状态，可选值如下： - normal：正常。 - soldout：售罄
	Phase map[string]string `json:"phase,omitempty"`
}

func (o ResourceFlavorStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceFlavorStatus struct{}"
	}

	return strings.Join([]string{"ResourceFlavorStatus", string(data)}, " ")
}
