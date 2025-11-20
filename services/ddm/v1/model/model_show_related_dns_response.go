package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRelatedDnsResponse Response Object
type ShowRelatedDnsResponse struct {

	// 关联DN。
	RelatedDataNodes *[]RelatedDnVo `json:"related_data_nodes,omitempty"`

	// 最近恢复时间点。
	LatestRestorableTime *string `json:"latest_restorable_time,omitempty"`
	HttpStatusCode       int     `json:"-"`
}

func (o ShowRelatedDnsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRelatedDnsResponse struct{}"
	}

	return strings.Join([]string{"ShowRelatedDnsResponse", string(data)}, " ")
}
