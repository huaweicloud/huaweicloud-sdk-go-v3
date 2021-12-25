package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Repository struct {
	// 创建仓库的UUID

	RepositoryUuid *string `json:"repository_uuid,omitempty"`
}

func (o Repository) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Repository struct{}"
	}

	return strings.Join([]string{"Repository", string(data)}, " ")
}
