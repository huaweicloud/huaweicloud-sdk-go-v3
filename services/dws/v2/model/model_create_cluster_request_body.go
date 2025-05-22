package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateClusterRequestBody **参数解释**： 创建集群的请求体。 **取值范围**： 非空对象。
type CreateClusterRequestBody struct {
	Cluster *CreateClusterInfo `json:"cluster"`
}

func (o CreateClusterRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateClusterRequestBody struct{}"
	}

	return strings.Join([]string{"CreateClusterRequestBody", string(data)}, " ")
}
