package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ApplyInternetRequest Request Object
type ApplyInternetRequest struct {
	Body *ApplyInternetReq `json:"body,omitempty"`
}

func (o ApplyInternetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplyInternetRequest struct{}"
	}

	return strings.Join([]string{"ApplyInternetRequest", string(data)}, " ")
}
