package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTtsJobResponse Response Object
type ShowTtsJobResponse struct {

	// 总记录数。
	Count *int32 `json:"count,omitempty"`

	// 语音合成任务列表。
	Data           *[]TtsJobDetail `json:"data,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ShowTtsJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTtsJobResponse struct{}"
	}

	return strings.Join([]string{"ShowTtsJobResponse", string(data)}, " ")
}
