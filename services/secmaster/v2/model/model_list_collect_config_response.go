package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCollectConfigResponse Response Object
type ListCollectConfigResponse struct {

	// 所有的云厂商、云产品和日志
	AllVendors *[]ListCollectConfigResponseBodyAllVendors `json:"all_vendors,omitempty"`

	CofingStatistics *ListCollectConfigResponseBodyCofingStatistics `json:"cofing_statistics,omitempty"`

	// 数据
	DataList *[]ListCollectConfigResponseBodyDataList `json:"data_list,omitempty"`

	// 数据集列表
	Datasets *[]DatasetInfo `json:"datasets,omitempty"`

	// 数据空间ID
	DataspaceId *string `json:"dataspace_id,omitempty"`

	// 数据空间名称
	DataspaceName *string `json:"dataspace_name,omitempty"`

	// 账号ID
	DomainId *string `json:"domain_id,omitempty"`

	// lts日志配置
	LtsSets *[]LtsResponseVo `json:"lts_sets,omitempty"`

	// 项目ID
	ProjectId *string `json:"project_id,omitempty"`

	// 所属region
	RegionId *string `json:"region_id,omitempty"`

	// 工作空间 id
	WorkspaceId    *string `json:"workspace_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListCollectConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCollectConfigResponse struct{}"
	}

	return strings.Join([]string{"ListCollectConfigResponse", string(data)}, " ")
}
