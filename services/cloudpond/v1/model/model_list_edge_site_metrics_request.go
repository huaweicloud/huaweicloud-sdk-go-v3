package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListEdgeSiteMetricsRequest Request Object
type ListEdgeSiteMetricsRequest struct {

	// 边缘小站ID
	SiteId string `json:"site_id"`

	// 指定维度查询 - site_id: 按站点维度，查询站点下计算、存储资源容量信息 - flavor: 按规格维度，查询站点下各flavor的计算资源使用情况 - storage: 按存储维度，查询站点下各存储资源类型的使用情况 - flavor_capacity: 按规格容量维度，查询站点下各规格可发放数量预测 - storage_pool：按存储池维度，查询站点下各存储池的使用情况
	Dim *ListEdgeSiteMetricsRequestDim `json:"dim,omitempty"`
}

func (o ListEdgeSiteMetricsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEdgeSiteMetricsRequest struct{}"
	}

	return strings.Join([]string{"ListEdgeSiteMetricsRequest", string(data)}, " ")
}

type ListEdgeSiteMetricsRequestDim struct {
	value string
}

type ListEdgeSiteMetricsRequestDimEnum struct {
	SITE_ID         ListEdgeSiteMetricsRequestDim
	FLAVOR          ListEdgeSiteMetricsRequestDim
	STORAGE         ListEdgeSiteMetricsRequestDim
	FLAVOR_CAPACITY ListEdgeSiteMetricsRequestDim
	STORAGE_POOL    ListEdgeSiteMetricsRequestDim
}

func GetListEdgeSiteMetricsRequestDimEnum() ListEdgeSiteMetricsRequestDimEnum {
	return ListEdgeSiteMetricsRequestDimEnum{
		SITE_ID: ListEdgeSiteMetricsRequestDim{
			value: "site_id",
		},
		FLAVOR: ListEdgeSiteMetricsRequestDim{
			value: "flavor",
		},
		STORAGE: ListEdgeSiteMetricsRequestDim{
			value: "storage",
		},
		FLAVOR_CAPACITY: ListEdgeSiteMetricsRequestDim{
			value: "flavor_capacity",
		},
		STORAGE_POOL: ListEdgeSiteMetricsRequestDim{
			value: "storage_pool",
		},
	}
}

func (c ListEdgeSiteMetricsRequestDim) Value() string {
	return c.value
}

func (c ListEdgeSiteMetricsRequestDim) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListEdgeSiteMetricsRequestDim) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
