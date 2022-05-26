package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ListDependenciesRequest struct {

	// 依赖包类型public：公开,private:私有，all：全部。缺省时查询全量。
	DependencyType *ListDependenciesRequestDependencyType `json:"dependency_type,omitempty"`

	// 运行时语言 Java11、Nodejs14.18、Python3.9在type为v2时支持
	Runtime *ListDependenciesRequestRuntime `json:"runtime,omitempty"`

	// 依赖包名称。
	Name *string `json:"name,omitempty"`

	// 上一次查询依赖包的最后记录位置，默认为\"0\"。
	Marker *string `json:"marker,omitempty"`

	// 本次查询可获取的依赖包的最大数目，默认为\"400\"。
	Limit *string `json:"limit,omitempty"`
}

func (o ListDependenciesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDependenciesRequest struct{}"
	}

	return strings.Join([]string{"ListDependenciesRequest", string(data)}, " ")
}

type ListDependenciesRequestDependencyType struct {
	value string
}

type ListDependenciesRequestDependencyTypeEnum struct {
	PUBLIC  ListDependenciesRequestDependencyType
	PRIVATE ListDependenciesRequestDependencyType
	ALL     ListDependenciesRequestDependencyType
}

func GetListDependenciesRequestDependencyTypeEnum() ListDependenciesRequestDependencyTypeEnum {
	return ListDependenciesRequestDependencyTypeEnum{
		PUBLIC: ListDependenciesRequestDependencyType{
			value: "public",
		},
		PRIVATE: ListDependenciesRequestDependencyType{
			value: "private",
		},
		ALL: ListDependenciesRequestDependencyType{
			value: "all",
		},
	}
}

func (c ListDependenciesRequestDependencyType) Value() string {
	return c.value
}

func (c ListDependenciesRequestDependencyType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListDependenciesRequestDependencyType) UnmarshalJSON(b []byte) error {
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

type ListDependenciesRequestRuntime struct {
	value string
}

type ListDependenciesRequestRuntimeEnum struct {
	JAVA8           ListDependenciesRequestRuntime
	JAVA11          ListDependenciesRequestRuntime
	NODE_JS6_10     ListDependenciesRequestRuntime
	NODE_JS8_10     ListDependenciesRequestRuntime
	NODE_JS10_16    ListDependenciesRequestRuntime
	NODE_JS12_13    ListDependenciesRequestRuntime
	NODE_JS14_18    ListDependenciesRequestRuntime
	PYTHON2_7       ListDependenciesRequestRuntime
	PYTHON3_6       ListDependenciesRequestRuntime
	PYTHON3_9       ListDependenciesRequestRuntime
	GO1_8           ListDependenciesRequestRuntime
	GO1_X           ListDependenciesRequestRuntime
	C__NET_CORE_2_0 ListDependenciesRequestRuntime
	C__NET_CORE_2_1 ListDependenciesRequestRuntime
	C__NET_CORE_3_1 ListDependenciesRequestRuntime
	PHP7_3          ListDependenciesRequestRuntime
}

func GetListDependenciesRequestRuntimeEnum() ListDependenciesRequestRuntimeEnum {
	return ListDependenciesRequestRuntimeEnum{
		JAVA8: ListDependenciesRequestRuntime{
			value: "Java8",
		},
		JAVA11: ListDependenciesRequestRuntime{
			value: "Java11",
		},
		NODE_JS6_10: ListDependenciesRequestRuntime{
			value: "Node.js6.10",
		},
		NODE_JS8_10: ListDependenciesRequestRuntime{
			value: "Node.js8.10",
		},
		NODE_JS10_16: ListDependenciesRequestRuntime{
			value: "Node.js10.16",
		},
		NODE_JS12_13: ListDependenciesRequestRuntime{
			value: "Node.js12.13",
		},
		NODE_JS14_18: ListDependenciesRequestRuntime{
			value: "Node.js14.18",
		},
		PYTHON2_7: ListDependenciesRequestRuntime{
			value: "Python2.7",
		},
		PYTHON3_6: ListDependenciesRequestRuntime{
			value: "Python3.6",
		},
		PYTHON3_9: ListDependenciesRequestRuntime{
			value: "Python3.9",
		},
		GO1_8: ListDependenciesRequestRuntime{
			value: "Go1.8",
		},
		GO1_X: ListDependenciesRequestRuntime{
			value: "Go1.x",
		},
		C__NET_CORE_2_0: ListDependenciesRequestRuntime{
			value: "C#(.NET Core 2.0)",
		},
		C__NET_CORE_2_1: ListDependenciesRequestRuntime{
			value: "C#(.NET Core 2.1)",
		},
		C__NET_CORE_3_1: ListDependenciesRequestRuntime{
			value: "C#(.NET Core 3.1)",
		},
		PHP7_3: ListDependenciesRequestRuntime{
			value: "PHP7.3",
		},
	}
}

func (c ListDependenciesRequestRuntime) Value() string {
	return c.value
}

func (c ListDependenciesRequestRuntime) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListDependenciesRequestRuntime) UnmarshalJSON(b []byte) error {
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
