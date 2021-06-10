package model

import (
	"encoding/json"

	"strings"
)

type UserInstantIncidentMsgV2 struct {
	// 工单id

	IncidentId *string `json:"incident_id,omitempty"`
	// 留言列表

	MessageList *[]QueryMessageInfoV2 `json:"message_list,omitempty"`
}

func (o UserInstantIncidentMsgV2) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UserInstantIncidentMsgV2 struct{}"
	}

	return strings.Join([]string{"UserInstantIncidentMsgV2", string(data)}, " ")
}
