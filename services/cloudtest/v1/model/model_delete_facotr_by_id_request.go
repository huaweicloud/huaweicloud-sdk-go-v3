package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteFacotrByIdRequest Request Object
type DeleteFacotrByIdRequest struct {

	// 项目ID，固定长度32位字符（字母和数字）。
	ProjectId string `json:"project_id"`

	// 因子ID
	Id string `json:"id"`
}

func (o DeleteFacotrByIdRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteFacotrByIdRequest struct{}"
	}

	return strings.Join([]string{"DeleteFacotrByIdRequest", string(data)}, " ")
}
