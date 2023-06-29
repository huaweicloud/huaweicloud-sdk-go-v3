package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DomainMigrate
type DomainMigrate struct {

	// 是否触发其他区域迁移
	AllRegions bool `json:"all_regions"`

	// 存储库默认扩容比，取值范围0到1
	Reservation float32 `json:"reservation"`
}

func (o DomainMigrate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DomainMigrate struct{}"
	}

	return strings.Join([]string{"DomainMigrate", string(data)}, " ")
}
