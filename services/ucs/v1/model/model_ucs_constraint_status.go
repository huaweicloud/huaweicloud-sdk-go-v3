package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UcsConstraintStatus struct {
	Violations *[]UcsStatusViolation `json:"violations,omitempty"`
}

func (o UcsConstraintStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UcsConstraintStatus struct{}"
	}

	return strings.Join([]string{"UcsConstraintStatus", string(data)}, " ")
}
