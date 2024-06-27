package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowFactorByIdRequest Request Object
type ShowFactorByIdRequest struct {

	// 项目ID，固定长度32位字符（字母和数字）。
	ProjectId string `json:"project_id"`

	// 因子ID
	Id string `json:"id"`
}

func (o ShowFactorByIdRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFactorByIdRequest struct{}"
	}

	return strings.Join([]string{"ShowFactorByIdRequest", string(data)}, " ")
}
