package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResizeStarRocksFlavorResponse Response Object
type ResizeStarRocksFlavorResponse struct {

	// 工作流ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ResizeStarRocksFlavorResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResizeStarRocksFlavorResponse struct{}"
	}

	return strings.Join([]string{"ResizeStarRocksFlavorResponse", string(data)}, " ")
}
