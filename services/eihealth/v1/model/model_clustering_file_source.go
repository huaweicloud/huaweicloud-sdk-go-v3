package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ClusteringFileSource 受体的数据源：用户私有数据中心、承载租户公共数据（含样例/公共库）。
type ClusteringFileSource struct {
	value string
}

type ClusteringFileSourceEnum struct {
	PRIVATE ClusteringFileSource
	PUBLIC  ClusteringFileSource
	RAW     ClusteringFileSource
}

func GetClusteringFileSourceEnum() ClusteringFileSourceEnum {
	return ClusteringFileSourceEnum{
		PRIVATE: ClusteringFileSource{
			value: "PRIVATE",
		},
		PUBLIC: ClusteringFileSource{
			value: "PUBLIC",
		},
		RAW: ClusteringFileSource{
			value: "RAW",
		},
	}
}

func (c ClusteringFileSource) Value() string {
	return c.value
}

func (c ClusteringFileSource) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ClusteringFileSource) UnmarshalJSON(b []byte) error {
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
