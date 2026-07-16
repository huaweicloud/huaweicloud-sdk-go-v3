package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NetworkMetadataUpdate 更新网络资源时的metadata信息。
type NetworkMetadataUpdate struct {
	Annotations *NetworkMetadataAnnotations `json:"annotations,omitempty"`
}

func (o NetworkMetadataUpdate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NetworkMetadataUpdate struct{}"
	}

	return strings.Join([]string{"NetworkMetadataUpdate", string(data)}, " ")
}
