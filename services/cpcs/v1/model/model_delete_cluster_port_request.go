package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DeleteClusterPortRequest Request Object
type DeleteClusterPortRequest struct {

	// 集群ID
	ClusterId string `json:"cluster_id"`

	// 端口的id
	Id string `json:"id"`

	// DELETE_DB 仅删除CPCS数据库中的记录，不删除elb资源。 TRY_DELETE 当被删除的端口的wrong字段为false时，删除elb资源。不为false时仅删除数据库。 FORCE_DELETE 删除cpcs数据库和elb资源
	Type DeleteClusterPortRequestType `json:"type"`
}

func (o DeleteClusterPortRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteClusterPortRequest struct{}"
	}

	return strings.Join([]string{"DeleteClusterPortRequest", string(data)}, " ")
}

type DeleteClusterPortRequestType struct {
	value string
}

type DeleteClusterPortRequestTypeEnum struct {
	DELETE_DB    DeleteClusterPortRequestType
	TRY_DELETE   DeleteClusterPortRequestType
	FORCE_DELETE DeleteClusterPortRequestType
}

func GetDeleteClusterPortRequestTypeEnum() DeleteClusterPortRequestTypeEnum {
	return DeleteClusterPortRequestTypeEnum{
		DELETE_DB: DeleteClusterPortRequestType{
			value: "DELETE_DB",
		},
		TRY_DELETE: DeleteClusterPortRequestType{
			value: "TRY_DELETE",
		},
		FORCE_DELETE: DeleteClusterPortRequestType{
			value: "FORCE_DELETE",
		},
	}
}

func (c DeleteClusterPortRequestType) Value() string {
	return c.value
}

func (c DeleteClusterPortRequestType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteClusterPortRequestType) UnmarshalJSON(b []byte) error {
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
