package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteGroupRequest Request Object
type DeleteGroupRequest struct {

	// CodeArts项目ID，32位数字、小写字母组合。
	ProjectId string `json:"project_id"`

	// 主键id
	Id string `json:"id"`
}

func (o DeleteGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteGroupRequest struct{}"
	}

	return strings.Join([]string{"DeleteGroupRequest", string(data)}, " ")
}
