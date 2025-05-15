package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AreaId Area ID。
type AreaId struct {
	AreaId *AreaIdDef `json:"area_id"`
}

func (o AreaId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AreaId struct{}"
	}

	return strings.Join([]string{"AreaId", string(data)}, " ")
}
