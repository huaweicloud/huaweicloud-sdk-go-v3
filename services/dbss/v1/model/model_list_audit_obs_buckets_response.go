package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAuditObsBucketsResponse Response Object
type ListAuditObsBucketsResponse struct {

	// OBS桶列表
	ObsList        *[]ObsBucketObject `json:"obs_list,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ListAuditObsBucketsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAuditObsBucketsResponse struct{}"
	}

	return strings.Join([]string{"ListAuditObsBucketsResponse", string(data)}, " ")
}
