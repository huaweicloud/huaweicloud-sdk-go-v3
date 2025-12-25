package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFilesRequest Request Object
type ListFilesRequest struct {
	Body *RepoFileQueryDtov5 `json:"body,omitempty"`
}

func (o ListFilesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFilesRequest struct{}"
	}

	return strings.Join([]string{"ListFilesRequest", string(data)}, " ")
}
