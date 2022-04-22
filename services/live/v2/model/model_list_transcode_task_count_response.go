package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListTranscodeTaskCountResponse struct {

	// 时间戳及相应时间的数值。
	TranscodeDataList *[]TranscodeCountData `json:"transcode_data_list,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListTranscodeTaskCountResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTranscodeTaskCountResponse struct{}"
	}

	return strings.Join([]string{"ListTranscodeTaskCountResponse", string(data)}, " ")
}
