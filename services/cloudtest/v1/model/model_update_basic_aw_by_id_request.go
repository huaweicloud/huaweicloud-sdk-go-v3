package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateBasicAwByIdRequest Request Object
type UpdateBasicAwByIdRequest struct {

	// 项目ID，固定长度32位字符（字母和数字）。
	ProjectId string `json:"project_id"`

	// AW的ID
	AwId string `json:"aw_id"`

	Body *UpdateBasicAwReq `json:"body,omitempty"`
}

func (o UpdateBasicAwByIdRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateBasicAwByIdRequest struct{}"
	}

	return strings.Join([]string{"UpdateBasicAwByIdRequest", string(data)}, " ")
}
