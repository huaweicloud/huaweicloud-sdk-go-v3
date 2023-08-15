package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExpandInstanceStorageRequest Request Object
type ExpandInstanceStorageRequest struct {

	// 集群的ID。
	ClusterId string `json:"cluster_id"`

	Body *ExpandInstanceStorage `json:"body,omitempty"`
}

func (o ExpandInstanceStorageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExpandInstanceStorageRequest struct{}"
	}

	return strings.Join([]string{"ExpandInstanceStorageRequest", string(data)}, " ")
}
