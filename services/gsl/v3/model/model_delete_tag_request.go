package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteTagRequest Request Object
type DeleteTagRequest struct {

	// 标签标识
	TagId int64 `json:"tag_id"`
}

func (o DeleteTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTagRequest struct{}"
	}

	return strings.Join([]string{"DeleteTagRequest", string(data)}, " ")
}
