package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ValidityRecurrence struct {
	Interval *ValidityInterval `json:"interval,omitempty"`

	Schedule *RecurrenceSchedule `json:"schedule,omitempty"`
}

func (o ValidityRecurrence) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ValidityRecurrence struct{}"
	}

	return strings.Join([]string{"ValidityRecurrence", string(data)}, " ")
}
