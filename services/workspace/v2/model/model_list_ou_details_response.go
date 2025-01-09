package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListOuDetailsResponse Response Object
type ListOuDetailsResponse struct {

	// OU对象
	Ous *[]OuNameInfo `json:"ous,omitempty"`

	// OU总记录数
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListOuDetailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListOuDetailsResponse struct{}"
	}

	return strings.Join([]string{"ListOuDetailsResponse", string(data)}, " ")
}
