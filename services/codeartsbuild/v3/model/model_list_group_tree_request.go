package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListGroupTreeRequest Request Object
type ListGroupTreeRequest struct {

	// CodeArts项目ID，32位数字、小写字母组合。
	ProjectId string `json:"project_id"`
}

func (o ListGroupTreeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGroupTreeRequest struct{}"
	}

	return strings.Join([]string{"ListGroupTreeRequest", string(data)}, " ")
}
