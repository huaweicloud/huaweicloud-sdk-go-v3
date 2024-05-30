package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TableModelStatisticVo struct {
	Total *StatisticStandardCoverageVo `json:"total,omitempty"`

	Published *StatisticStandardCoverageVo `json:"published,omitempty"`

	SyncTotal *SyncStatusStatisticVo `json:"sync_total,omitempty"`

	PhysicalTable *SyncStatusStatisticVo `json:"physical_table,omitempty"`

	TechnicalAsset *SyncStatusStatisticVo `json:"technical_asset,omitempty"`

	BusinessAsset *SyncStatusStatisticVo `json:"business_asset,omitempty"`

	MetaDataLink *SyncStatusStatisticVo `json:"meta_data_link,omitempty"`

	DataQuality *SyncStatusStatisticVo `json:"data_quality,omitempty"`
}

func (o TableModelStatisticVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TableModelStatisticVo struct{}"
	}

	return strings.Join([]string{"TableModelStatisticVo", string(data)}, " ")
}
