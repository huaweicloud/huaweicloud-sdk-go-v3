package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateChInstanceRequestBodyVolume 存储规格。
type CreateChInstanceRequestBodyVolume struct {

	// 磁盘IO类型。 取值范围： - SSD：超高IO - ESSD：极速型SSD
	IoType string `json:"io_type"`

	// 磁盘容量。取值范围：50GB~32000GB。
	CapacityInGb int32 `json:"capacity_in_gb"`
}

func (o CreateChInstanceRequestBodyVolume) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateChInstanceRequestBodyVolume struct{}"
	}

	return strings.Join([]string{"CreateChInstanceRequestBodyVolume", string(data)}, " ")
}
