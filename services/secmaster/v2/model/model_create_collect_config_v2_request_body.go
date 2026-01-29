package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateCollectConfigV2RequestBody struct {

	// 数据集列表
	Config []ConfigInfo `json:"config"`

	// 数据空间ID
	DataspaceId string `json:"dataspace_id"`

	// 数据空间名称
	DataspaceName string `json:"dataspace_name"`

	// 账号ID
	DomainId *string `json:"domain_id,omitempty"`

	// lts配置
	LtsConfig *[]LtsRquestVo `json:"lts_config,omitempty"`

	// 工作空间 id
	WorkspaceId string `json:"workspace_id"`
}

func (o CreateCollectConfigV2RequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCollectConfigV2RequestBody struct{}"
	}

	return strings.Join([]string{"CreateCollectConfigV2RequestBody", string(data)}, " ")
}
