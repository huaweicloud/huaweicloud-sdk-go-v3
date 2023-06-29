package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowJobDetailResponse Response Object
type ShowJobDetailResponse struct {
	Job            *JobDetailResp `json:"job,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ShowJobDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowJobDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowJobDetailResponse", string(data)}, " ")
}
