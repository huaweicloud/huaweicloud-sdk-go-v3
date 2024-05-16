package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateBusinessResponse Response Object
type CreateBusinessResponse struct {

	// 应用的id
	Id *int32 `json:"id,omitempty"`

	Ok             *string `json:"ok,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateBusinessResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateBusinessResponse struct{}"
	}

	return strings.Join([]string{"CreateBusinessResponse", string(data)}, " ")
}
