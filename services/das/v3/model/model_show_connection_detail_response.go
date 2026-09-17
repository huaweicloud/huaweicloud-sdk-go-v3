package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowConnectionDetailResponse Response Object
type ShowConnectionDetailResponse struct {
	DasConnInfo    *DasConnInfo `json:"das_conn_info,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ShowConnectionDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowConnectionDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowConnectionDetailResponse", string(data)}, " ")
}
