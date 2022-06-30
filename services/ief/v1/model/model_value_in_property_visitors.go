package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 孪生属性配置，与access_protocol关联。
type ValueInPropertyVisitors struct {
	RegisterType *ValueInAttributes `json:"register_type,omitempty"`

	AccessMode *ValueInAttributes `json:"access_mode,omitempty"`

	RegisterIndex *ValueInAttributes `json:"register_index,omitempty"`

	RegisterNum *ValueInAttributes `json:"register_num,omitempty"`

	ScaleIndex *ValueInAttributes `json:"scale_index,omitempty"`

	OriginalDatatype *ValueInAttributes `json:"original_datatype,omitempty"`

	ExpectedDatatype *ValueInAttributes `json:"expected_datatype,omitempty"`

	IsRegisterswap *ValueInAttributes `json:"is_registerswap,omitempty"`

	IsSwap *ValueInAttributes `json:"is_swap,omitempty"`

	SampleInterval *ValueInAttributes `json:"sample_interval,omitempty"`

	DataMin *ValueInAttributes `json:"data_min,omitempty"`

	DataMax *ValueInAttributes `json:"data_max,omitempty"`

	NodeId *ValueInAttributes `json:"node_id,omitempty"`

	BrowseName *ValueInAttributes `json:"browse_name,omitempty"`
}

func (o ValueInPropertyVisitors) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ValueInPropertyVisitors struct{}"
	}

	return strings.Join([]string{"ValueInPropertyVisitors", string(data)}, " ")
}
