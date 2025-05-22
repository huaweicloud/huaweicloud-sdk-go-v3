package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Cluster **参数解释**： 集群对象信息。 **取值范围**： 不涉及。
type Cluster struct {

	// **参数解释**： 集群ID。 **取值范围**： 不涉及。
	Id string `json:"id"`
}

func (o Cluster) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Cluster struct{}"
	}

	return strings.Join([]string{"Cluster", string(data)}, " ")
}
