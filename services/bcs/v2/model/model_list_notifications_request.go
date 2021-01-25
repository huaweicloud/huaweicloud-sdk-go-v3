package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListNotificationsRequest struct {
}

func (o ListNotificationsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListNotificationsRequest struct{}"
	}

	return strings.Join([]string{"ListNotificationsRequest", string(data)}, " ")
}
