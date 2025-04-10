package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCrossCloudDisasterRelationsResponse Response Object
type ShowCrossCloudDisasterRelationsResponse struct {

	// 查询结果总数。
	Total *int32 `json:"total,omitempty"`

	// 容灾关系详情。
	Relations      *[]DisasterRelations `json:"relations,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ShowCrossCloudDisasterRelationsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCrossCloudDisasterRelationsResponse struct{}"
	}

	return strings.Join([]string{"ShowCrossCloudDisasterRelationsResponse", string(data)}, " ")
}
