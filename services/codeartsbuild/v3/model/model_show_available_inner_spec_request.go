package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAvailableInnerSpecRequest Request Object
type ShowAvailableInnerSpecRequest struct {

	// CodeArts项目ID，32位数字、小写字母组合。
	ProjectId string `json:"project_id"`

	// 使用机器的架构，如x86-64、arm等
	Arch string `json:"arch"`
}

func (o ShowAvailableInnerSpecRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAvailableInnerSpecRequest struct{}"
	}

	return strings.Join([]string{"ShowAvailableInnerSpecRequest", string(data)}, " ")
}
