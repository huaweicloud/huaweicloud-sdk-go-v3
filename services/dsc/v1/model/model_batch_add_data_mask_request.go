package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type BatchAddDataMaskRequest struct {
	Body *DynamicDataMask `json:"body,omitempty" xml:"body"`
}

func (o BatchAddDataMaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAddDataMaskRequest struct{}"
	}

	return strings.Join([]string{"BatchAddDataMaskRequest", string(data)}, " ")
}
