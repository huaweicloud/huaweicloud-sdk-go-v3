package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFileBlameLinesResponse Response Object
type ListFileBlameLinesResponse struct {

	// 溯源信息
	Body           *[]BlameDto `json:"body,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ListFileBlameLinesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFileBlameLinesResponse struct{}"
	}

	return strings.Join([]string{"ListFileBlameLinesResponse", string(data)}, " ")
}
