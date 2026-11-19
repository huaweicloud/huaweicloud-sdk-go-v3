package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SupabaseSpecFlavorInfoAvailabilityZones **约束限制**  不涉及  **取值范围**  不涉及。  **默认取值**  不涉及。
type SupabaseSpecFlavorInfoAvailabilityZones struct {
	Code *string `json:"code,omitempty"`

	Status *string `json:"status,omitempty"`
}

func (o SupabaseSpecFlavorInfoAvailabilityZones) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SupabaseSpecFlavorInfoAvailabilityZones struct{}"
	}

	return strings.Join([]string{"SupabaseSpecFlavorInfoAvailabilityZones", string(data)}, " ")
}
