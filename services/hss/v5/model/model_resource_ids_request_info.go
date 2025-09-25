package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResourceIdsRequestInfo struct {

	// **参数解释**: 资源ID列表 **约束限制**: 不涉及 **取值范围**: 不涉及 **默认取值**: 不涉及
	ResourceIdList []string `json:"resource_id_list"`
}

func (o ResourceIdsRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceIdsRequestInfo struct{}"
	}

	return strings.Join([]string{"ResourceIdsRequestInfo", string(data)}, " ")
}
