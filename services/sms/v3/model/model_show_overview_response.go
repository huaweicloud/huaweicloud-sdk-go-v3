package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowOverviewResponse struct {

	// 等待中
	Waiting *int32 `json:"waiting,omitempty" xml:"waiting"`

	// 复制中
	Replicating *int32 `json:"replicating,omitempty" xml:"replicating"`

	// 同步中
	Syncing *int32 `json:"syncing,omitempty" xml:"syncing"`

	// 其它
	Other          *int32 `json:"other,omitempty" xml:"other"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowOverviewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowOverviewResponse struct{}"
	}

	return strings.Join([]string{"ShowOverviewResponse", string(data)}, " ")
}
