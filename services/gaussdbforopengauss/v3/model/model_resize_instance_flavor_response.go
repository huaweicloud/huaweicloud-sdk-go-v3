package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ResizeInstanceFlavorResponse struct {

	// 任务id。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ResizeInstanceFlavorResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResizeInstanceFlavorResponse struct{}"
	}

	return strings.Join([]string{"ResizeInstanceFlavorResponse", string(data)}, " ")
}
