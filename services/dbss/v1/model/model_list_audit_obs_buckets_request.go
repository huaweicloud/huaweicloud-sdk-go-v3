package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAuditObsBucketsRequest Request Object
type ListAuditObsBucketsRequest struct {
}

func (o ListAuditObsBucketsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAuditObsBucketsRequest struct{}"
	}

	return strings.Join([]string{"ListAuditObsBucketsRequest", string(data)}, " ")
}
