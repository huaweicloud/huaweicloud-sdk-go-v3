package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ApplyDesktopsInternetRequest Request Object
type ApplyDesktopsInternetRequest struct {
	Body *ApplyDesktopsInternetReq `json:"body,omitempty"`
}

func (o ApplyDesktopsInternetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplyDesktopsInternetRequest struct{}"
	}

	return strings.Join([]string{"ApplyDesktopsInternetRequest", string(data)}, " ")
}
