package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListVersionsResponse Response Object
type ListVersionsResponse struct {

	// 版本信息。
	Versions       *[]Versions `json:"versions,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ListVersionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVersionsResponse struct{}"
	}

	return strings.Join([]string{"ListVersionsResponse", string(data)}, " ")
}
