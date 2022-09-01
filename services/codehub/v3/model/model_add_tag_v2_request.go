package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type AddTagV2Request struct {

	// 仓库主键id
	RepositoryId int32 `json:"repository_id" xml:"repository_id"`

	Body *AddTagsRequest `json:"body,omitempty" xml:"body"`
}

func (o AddTagV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddTagV2Request struct{}"
	}

	return strings.Join([]string{"AddTagV2Request", string(data)}, " ")
}
