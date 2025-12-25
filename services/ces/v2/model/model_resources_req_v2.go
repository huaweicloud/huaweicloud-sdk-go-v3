package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResourcesReqV2 struct {

	// **参数解释**： 资源信息列表。 **约束限制**： 包含的资源信息最多为1000个，最少为0个。
	Resources [][]Dimension `json:"resources"`
}

func (o ResourcesReqV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourcesReqV2 struct{}"
	}

	return strings.Join([]string{"ResourcesReqV2", string(data)}, " ")
}
