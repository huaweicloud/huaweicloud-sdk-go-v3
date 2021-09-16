package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type BatchAddDataMaskRequest struct {
	Body *DynamicDataMask `json:"body,omitempty"`
}

func (o BatchAddDataMaskRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchAddDataMaskRequest struct{}"
	}

	return strings.Join([]string{"BatchAddDataMaskRequest", string(data)}, " ")
}
