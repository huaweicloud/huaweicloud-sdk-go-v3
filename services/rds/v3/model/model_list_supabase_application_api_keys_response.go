package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSupabaseApplicationApiKeysResponse Response Object
type ListSupabaseApplicationApiKeysResponse struct {
	TotalCount *int64 `json:"total_count,omitempty"`

	Data           *[]SupabaseApplicationApiKeyRecord `json:"data,omitempty"`
	HttpStatusCode int                                `json:"-"`
}

func (o ListSupabaseApplicationApiKeysResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSupabaseApplicationApiKeysResponse struct{}"
	}

	return strings.Join([]string{"ListSupabaseApplicationApiKeysResponse", string(data)}, " ")
}
