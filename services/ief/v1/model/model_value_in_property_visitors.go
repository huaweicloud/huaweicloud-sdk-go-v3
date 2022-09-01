package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 孪生属性配置，与access_protocol关联。
type ValueInPropertyVisitors struct {
	RegisterType *ValueInAttributes `json:"register_type,omitempty" xml:"register_type"`

	AccessMode *ValueInAttributes `json:"access_mode,omitempty" xml:"access_mode"`

	RegisterIndex *ValueInAttributes `json:"register_index,omitempty" xml:"register_index"`

	RegisterNum *ValueInAttributes `json:"register_num,omitempty" xml:"register_num"`

	ScaleIndex *ValueInAttributes `json:"scale_index,omitempty" xml:"scale_index"`

	OriginalDatatype *ValueInAttributes `json:"original_datatype,omitempty" xml:"original_datatype"`

	ExpectedDatatype *ValueInAttributes `json:"expected_datatype,omitempty" xml:"expected_datatype"`

	IsRegisterswap *ValueInAttributes `json:"is_registerswap,omitempty" xml:"is_registerswap"`

	IsSwap *ValueInAttributes `json:"is_swap,omitempty" xml:"is_swap"`

	SampleInterval *ValueInAttributes `json:"sample_interval,omitempty" xml:"sample_interval"`

	DataMin *ValueInAttributes `json:"data_min,omitempty" xml:"data_min"`

	DataMax *ValueInAttributes `json:"data_max,omitempty" xml:"data_max"`

	NodeId *ValueInAttributes `json:"node_id,omitempty" xml:"node_id"`

	BrowseName *ValueInAttributes `json:"browse_name,omitempty" xml:"browse_name"`
}

func (o ValueInPropertyVisitors) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ValueInPropertyVisitors struct{}"
	}

	return strings.Join([]string{"ValueInPropertyVisitors", string(data)}, " ")
}
