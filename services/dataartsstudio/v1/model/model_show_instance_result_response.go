package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInstanceResultResponse Response Object
type ShowInstanceResultResponse struct {

	// 总数量
	Count *int32 `json:"count,omitempty"`

	// resources
	Resources      *[]SubInstanceResult `json:"resources,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ShowInstanceResultResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceResultResponse struct{}"
	}

	return strings.Join([]string{"ShowInstanceResultResponse", string(data)}, " ")
}
