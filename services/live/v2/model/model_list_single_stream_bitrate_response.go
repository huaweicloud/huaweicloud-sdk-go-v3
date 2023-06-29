package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSingleStreamBitrateResponse Response Object
type ListSingleStreamBitrateResponse struct {

	// 用量详情。
	BitrateInfoList *[]V2BitrateInfo `json:"bitrate_info_list,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListSingleStreamBitrateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSingleStreamBitrateResponse struct{}"
	}

	return strings.Join([]string{"ListSingleStreamBitrateResponse", string(data)}, " ")
}
