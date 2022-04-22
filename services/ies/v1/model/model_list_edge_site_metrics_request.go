package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ListEdgeSiteMetricsRequest struct {

	// 边缘小站ID
	SiteId string `json:"site_id"`

	// 指定维度查询 - site_id: 按站点维度，查询站点下计算、存储资源容量信息 - flavor: 按规格维度，查询站点下各flavor的计算资源使用情况
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
	SITE_ID ListEdgeSiteMetricsRequestDim
	FLAVOR  ListEdgeSiteMetricsRequestDim
}

func GetListEdgeSiteMetricsRequestDimEnum() ListEdgeSiteMetricsRequestDimEnum {
	return ListEdgeSiteMetricsRequestDimEnum{
		SITE_ID: ListEdgeSiteMetricsRequestDim{
			value: "site_id",
		},
		FLAVOR: ListEdgeSiteMetricsRequestDim{
			value: "flavor",
		},
	}
}

func (c ListEdgeSiteMetricsRequestDim) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListEdgeSiteMetricsRequestDim) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
