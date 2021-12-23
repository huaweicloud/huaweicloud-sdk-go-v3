package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type BatchAddDataMaskRequest struct {
	Body *DynamicDataMask `json:"body,omitempty"`
}

func (o BatchAddDataMaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAddDataMaskRequest struct{}"
	}

	return strings.Join([]string{"BatchAddDataMaskRequest", string(data)}, " ")
}
