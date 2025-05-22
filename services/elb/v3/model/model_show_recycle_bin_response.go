package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRecycleBinResponse Response Object
type ShowRecycleBinResponse struct {
	RecycleBin     *RecycleBinResponseBody `json:"recycle_bin,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o ShowRecycleBinResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRecycleBinResponse struct{}"
	}

	return strings.Join([]string{"ShowRecycleBinResponse", string(data)}, " ")
}
