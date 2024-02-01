package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type GcbBindingServiceAll struct {

	// 功能说明：绑定的服务类型。实例类型： - CC: 云连接 - GEIP: 全域弹性公网IP - GCN: 中心网络 - GSN: 分支网络 - ALL: 所有实例类型
	BindingService *GcbBindingServiceAllBindingService `json:"binding_service,omitempty"`
}

func (o GcbBindingServiceAll) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GcbBindingServiceAll struct{}"
	}

	return strings.Join([]string{"GcbBindingServiceAll", string(data)}, " ")
}

type GcbBindingServiceAllBindingService struct {
	value string
}

type GcbBindingServiceAllBindingServiceEnum struct {
	CC   GcbBindingServiceAllBindingService
	GEIP GcbBindingServiceAllBindingService
	GCN  GcbBindingServiceAllBindingService
	GSN  GcbBindingServiceAllBindingService
	ALL  GcbBindingServiceAllBindingService
}

func GetGcbBindingServiceAllBindingServiceEnum() GcbBindingServiceAllBindingServiceEnum {
	return GcbBindingServiceAllBindingServiceEnum{
		CC: GcbBindingServiceAllBindingService{
			value: "CC",
		},
		GEIP: GcbBindingServiceAllBindingService{
			value: "GEIP",
		},
		GCN: GcbBindingServiceAllBindingService{
			value: "GCN",
		},
		GSN: GcbBindingServiceAllBindingService{
			value: "GSN",
		},
		ALL: GcbBindingServiceAllBindingService{
			value: "ALL",
		},
	}
}

func (c GcbBindingServiceAllBindingService) Value() string {
	return c.value
}

func (c GcbBindingServiceAllBindingService) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GcbBindingServiceAllBindingService) UnmarshalJSON(b []byte) error {
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
