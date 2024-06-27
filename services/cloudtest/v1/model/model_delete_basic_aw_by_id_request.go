package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteBasicAwByIdRequest Request Object
type DeleteBasicAwByIdRequest struct {

	// 项目ID，固定长度32位字符（字母和数字）。
	ProjectId string `json:"project_id"`

	// AW ID
	AwId string `json:"aw_id"`
}

func (o DeleteBasicAwByIdRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteBasicAwByIdRequest struct{}"
	}

	return strings.Join([]string{"DeleteBasicAwByIdRequest", string(data)}, " ")
}
