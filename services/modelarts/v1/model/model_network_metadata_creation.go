package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NetworkMetadataCreation 网络资源metadata信息创建请求体。
type NetworkMetadataCreation struct {
	Labels *NetworkMetadataLabels `json:"labels"`
}

func (o NetworkMetadataCreation) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NetworkMetadataCreation struct{}"
	}

	return strings.Join([]string{"NetworkMetadataCreation", string(data)}, " ")
}
