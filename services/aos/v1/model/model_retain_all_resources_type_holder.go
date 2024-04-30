package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RetainAllResourcesTypeHolder struct {

	// 删除资源栈是否保留资源的标志位，如果不传默认为false，即默认不保留资源（删除资源栈后会删除资源栈中的资源）  * DeleteStackEnhanced API中，如果该参数未在RequestBody中给予，则删除时不会保留资源栈中的资源*
	RetainAllResources *bool `json:"retain_all_resources,omitempty"`
}

func (o RetainAllResourcesTypeHolder) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RetainAllResourcesTypeHolder struct{}"
	}

	return strings.Join([]string{"RetainAllResourcesTypeHolder", string(data)}, " ")
}
