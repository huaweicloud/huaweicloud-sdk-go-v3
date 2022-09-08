package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 服务详情
type ServiceRespDetail struct {

	// 创建时间
	CreatedAt string `json:"created_at"`

	// 服务ID
	Id string `json:"id"`

	MetaData *SvcMetadata `json:"meta_data"`

	// 租户ID
	ProjectId string `json:"project_id"`

	Spec *SvcSpec `json:"spec"`

	// 更新时间
	UpdatedAt string `json:"updated_at"`
}

func (o ServiceRespDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServiceRespDetail struct{}"
	}

	return strings.Join([]string{"ServiceRespDetail", string(data)}, " ")
}
