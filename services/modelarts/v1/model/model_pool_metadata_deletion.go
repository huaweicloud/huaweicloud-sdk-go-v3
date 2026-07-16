package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PoolMetadataDeletion 删除资源池时的metadata信息。
type PoolMetadataDeletion struct {

	// **参数描述**： 系统自动生成的pool名称，相当于poolId。 **取值范围**： 不涉及。
	Name string `json:"name"`

	// **参数描述**： 时间戳，例如\"2021-11-01T03:49:41Z\"。 **取值范围**： 不涉及。
	CreationTimestamp string `json:"creationTimestamp"`

	// **参数描述**： 时间戳，例如\"2021-11-01T03:49:41Z\"。 **取值范围**： 不涉及。
	DeletionTimestamp string `json:"deletionTimestamp"`

	Labels *PoolMetaLabels `json:"labels"`

	Annotations *PoolMetaAnnotations `json:"annotations,omitempty"`
}

func (o PoolMetadataDeletion) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoolMetadataDeletion struct{}"
	}

	return strings.Join([]string{"PoolMetadataDeletion", string(data)}, " ")
}
