package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListApiVersionsResponse struct {
	// 描述version相关对象的列表

	Versions       *[]VersionDetail `json:"versions,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ListApiVersionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListApiVersionsResponse struct{}"
	}

	return strings.Join([]string{"ListApiVersionsResponse", string(data)}, " ")
}
