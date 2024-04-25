package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSecuritySensitiveDataOverviewsResponse Response Object
type ListSecuritySensitiveDataOverviewsResponse struct {

	// 基于密级的概览统计
	SecrecyLevelStatistics *[]SensitiveDataSecrecyLevelOverviewQueryDto `json:"secrecy_level_statistics,omitempty"`

	// 基于分类的概览统计
	CategoryStatistics *[]SensitiveDataCategoryOverviewQueryDto `json:"category_statistics,omitempty"`
	HttpStatusCode     int                                      `json:"-"`
}

func (o ListSecuritySensitiveDataOverviewsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSecuritySensitiveDataOverviewsResponse struct{}"
	}

	return strings.Join([]string{"ListSecuritySensitiveDataOverviewsResponse", string(data)}, " ")
}
