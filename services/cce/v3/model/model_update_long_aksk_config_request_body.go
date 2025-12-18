package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateLongAkskConfigRequestBody struct {

	// **参数解释：** 新建集群是否启用LongAKSK。 **约束限制：** 不涉及 **取值范围：** - false: 新建集群不启用LongAKSK - true: 新建集群启用LongAKSK  **默认取值：** 不涉及
	EnableLongAKSKInNewCluster bool `json:"enableLongAKSKInNewCluster"`
}

func (o UpdateLongAkskConfigRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLongAkskConfigRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateLongAkskConfigRequestBody", string(data)}, " ")
}
