package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssociationAlarmTotal 告警模板关联的告警规则数目
type AssociationAlarmTotal struct {
}

func (o AssociationAlarmTotal) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociationAlarmTotal struct{}"
	}

	return strings.Join([]string{"AssociationAlarmTotal", string(data)}, " ")
}
