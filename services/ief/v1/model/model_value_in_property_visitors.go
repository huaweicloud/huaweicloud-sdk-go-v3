package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 孪生属性配置，与access_protocol关联。
type ValueInPropertyVisitors struct {
	RegisterType *ValueInPropertyVisitorsRegisterType `json:"register_type,omitempty"`

	AccessMode *ValueInPropertyVisitorsAccessMode `json:"access_mode,omitempty"`

	RegisterIndex *ValueInPropertyVisitorsRegisterIndex `json:"register_index,omitempty"`

	RegisterNum *ValueInPropertyVisitorsRegisterNum `json:"register_num,omitempty"`

	ScaleIndex *ValueInPropertyVisitorsScaleIndex `json:"scale_index,omitempty"`

	OriginalDatatype *ValueInPropertyVisitorsOriginalDatatype `json:"original_datatype,omitempty"`

	ExpectedDatatype *ValueInPropertyVisitorsExpectedDatatype `json:"expected_datatype,omitempty"`

	IsRegisterswap *ValueInPropertyVisitorsIsRegisterswap `json:"is_registerswap,omitempty"`

	IsSwap *ValueInPropertyVisitorsIsSwap `json:"is_swap,omitempty"`

	SampleInterval *ValueInPropertyVisitorsSampleInterval `json:"sample_interval,omitempty"`

	DataMin *ValueInPropertyVisitorsDataMin `json:"data_min,omitempty"`

	DataMax *ValueInPropertyVisitorsDataMax `json:"data_max,omitempty"`

	NodeId *ValueInPropertyVisitorsNodeId `json:"node_id,omitempty"`

	BrowseName *ValueInPropertyVisitorsBrowseName `json:"browse_name,omitempty"`
}

func (o ValueInPropertyVisitors) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ValueInPropertyVisitors struct{}"
	}

	return strings.Join([]string{"ValueInPropertyVisitors", string(data)}, " ")
}
