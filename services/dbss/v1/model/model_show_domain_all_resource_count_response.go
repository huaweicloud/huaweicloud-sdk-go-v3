package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDomainAllResourceCountResponse Response Object
type ShowDomainAllResourceCountResponse struct {

	// 总量
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowDomainAllResourceCountResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDomainAllResourceCountResponse struct{}"
	}

	return strings.Join([]string{"ShowDomainAllResourceCountResponse", string(data)}, " ")
}
