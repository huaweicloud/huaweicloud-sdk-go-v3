package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ValidateDeletableReplicaResponse Response Object
type ValidateDeletableReplicaResponse struct {

	// 是否有可供选择的副本组进行删除。
	CheckResult *bool `json:"check_result,omitempty"`

	// 可选的可用区ID列表
	AvailableZone *string `json:"available_zone,omitempty"`

	// 可选的保留节点列表
	ReplicationList *[]ReplicationInfo `json:"replication_list,omitempty"`
	HttpStatusCode  int                `json:"-"`
}

func (o ValidateDeletableReplicaResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ValidateDeletableReplicaResponse struct{}"
	}

	return strings.Join([]string{"ValidateDeletableReplicaResponse", string(data)}, " ")
}
