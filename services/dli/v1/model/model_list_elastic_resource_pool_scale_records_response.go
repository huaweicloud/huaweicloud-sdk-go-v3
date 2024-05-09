package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListElasticResourcePoolScaleRecordsResponse Response Object
type ListElasticResourcePoolScaleRecordsResponse struct {

	// 返回数组长度
	Count *int64 `json:"count,omitempty"`

	// 数组中返回的数据
	Items *[]ElasticResourcePoolScaleRecord `json:"items,omitempty"`

	XAuthToken     *string `json:"X-Auth-Token,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListElasticResourcePoolScaleRecordsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListElasticResourcePoolScaleRecordsResponse struct{}"
	}

	return strings.Join([]string{"ListElasticResourcePoolScaleRecordsResponse", string(data)}, " ")
}
