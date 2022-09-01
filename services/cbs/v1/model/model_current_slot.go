package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type CurrentSlot struct {

	// 槽位ID。
	SlotId string `json:"slot_id" xml:"slot_id"`

	// 槽位名称。
	SlotName string `json:"slot_name" xml:"slot_name"`

	// 槽位值。
	SlotValues []SlotValue `json:"slot_values" xml:"slot_values"`

	// 槽位标识。
	SlotIdentification *string `json:"slot_identification,omitempty" xml:"slot_identification"`
}

func (o CurrentSlot) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CurrentSlot struct{}"
	}

	return strings.Join([]string{"CurrentSlot", string(data)}, " ")
}
