package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ApplyDesktopsInternetResponse Response Object
type ApplyDesktopsInternetResponse struct {

	// 任务id。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ApplyDesktopsInternetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplyDesktopsInternetResponse struct{}"
	}

	return strings.Join([]string{"ApplyDesktopsInternetResponse", string(data)}, " ")
}
