package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BaseMultiActiveZoneSpec struct {
}

func (o BaseMultiActiveZoneSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BaseMultiActiveZoneSpec struct{}"
	}

	return strings.Join([]string{"BaseMultiActiveZoneSpec", string(data)}, " ")
}
