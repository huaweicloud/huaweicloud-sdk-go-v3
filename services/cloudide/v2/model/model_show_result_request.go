package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowResultRequest Request Object
type ShowResultRequest struct {

	// the unique if of the request
	RequestId string `json:"request_id"`
}

func (o ShowResultRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowResultRequest struct{}"
	}

	return strings.Join([]string{"ShowResultRequest", string(data)}, " ")
}
