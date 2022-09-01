package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListSharesResponse struct {

	// SFS Turbo文件系统的列表。
	Shares *[]Shares `json:"shares,omitempty" xml:"shares"`

	// SFS Turbo文件系统的数量。
	Count          *int32 `json:"count,omitempty" xml:"count"`
	HttpStatusCode int    `json:"-"`
}

func (o ListSharesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSharesResponse struct{}"
	}

	return strings.Join([]string{"ListSharesResponse", string(data)}, " ")
}
