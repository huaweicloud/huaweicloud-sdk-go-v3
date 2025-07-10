package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAdOusResponse Response Object
type ListAdOusResponse struct {

	// OU对象。
	OuInfos *[]AdOuInfo `json:"ou_infos,omitempty"`

	// OU总记录数。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListAdOusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAdOusResponse struct{}"
	}

	return strings.Join([]string{"ListAdOusResponse", string(data)}, " ")
}
