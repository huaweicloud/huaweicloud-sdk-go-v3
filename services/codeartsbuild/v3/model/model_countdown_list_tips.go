package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CountdownListTips 咨询
type CountdownListTips struct {
	Summary *Tips `json:"summary,omitempty"`
}

func (o CountdownListTips) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountdownListTips struct{}"
	}

	return strings.Join([]string{"CountdownListTips", string(data)}, " ")
}
