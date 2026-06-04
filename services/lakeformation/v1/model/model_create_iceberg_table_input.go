package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateIcebergTableInput struct {
	Schema *IcebergSchema `json:"schema"`

	PartitionSpec *IcebergPartitionSpec `json:"partition_spec,omitempty"`

	WriteOrder *IcebergSortOrder `json:"write_order,omitempty"`
}

func (o CreateIcebergTableInput) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateIcebergTableInput struct{}"
	}

	return strings.Join([]string{"CreateIcebergTableInput", string(data)}, " ")
}
