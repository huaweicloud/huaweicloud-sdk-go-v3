package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckRepositoryDuplicateNameRequest Request Object
type CheckRepositoryDuplicateNameRequest struct {

	// 语言类型 中文:zh-cn 英文:en-us
	XLanguage *string `json:"X-Language,omitempty"`

	// 项目id
	ProjectId string `json:"project_id"`

	// 仓库名称
	Name string `json:"name"`

	// 区域id
	RegionId string `json:"region_id"`
}

func (o CheckRepositoryDuplicateNameRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckRepositoryDuplicateNameRequest struct{}"
	}

	return strings.Join([]string{"CheckRepositoryDuplicateNameRequest", string(data)}, " ")
}
