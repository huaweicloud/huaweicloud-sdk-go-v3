package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CountCertResourceInstancesResponse Response Object
type CountCertResourceInstancesResponse struct {

	// 总记录数。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o CountCertResourceInstancesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountCertResourceInstancesResponse struct{}"
	}

	return strings.Join([]string{"CountCertResourceInstancesResponse", string(data)}, " ")
}
