package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MetaStoreId **参数解释**：LakeFormation服务的实例ID，可在LakeFormation控制台实例列表中查看，或通过调用LakeFormation服务的查询实例列表接口获取。 **约束限制**：不涉及。 **取值范围**：长度为32~36的英文字母、数字和中划线的组合。 **默认取值**：不涉及。
type MetaStoreId struct {
}

func (o MetaStoreId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MetaStoreId struct{}"
	}

	return strings.Join([]string{"MetaStoreId", string(data)}, " ")
}
