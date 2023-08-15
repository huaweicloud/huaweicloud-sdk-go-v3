package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ImportPointsRequest Request Object
type ImportPointsRequest struct {

	// 边缘节点ID
	EdgeNodeId string `json:"edge_node_id"`

	// 采集数据源id，创建数据源配置时设置，节点下唯一。
	DsId string `json:"ds_id"`

	// 该字段PARTIAL则增量覆盖，已有点位更新，新增点位插入;该字段为COMPLETE则全量覆盖，则删除数据源下所有点位，插入当前导入所有点位
	UpdateType ImportPointsRequestUpdateType `json:"update_type"`

	Body *ImportPointsRequestBody `json:"body,omitempty" type:"multipart"`
}

func (o ImportPointsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportPointsRequest struct{}"
	}

	return strings.Join([]string{"ImportPointsRequest", string(data)}, " ")
}

type ImportPointsRequestUpdateType struct {
	value string
}

type ImportPointsRequestUpdateTypeEnum struct {
	PARTIAL  ImportPointsRequestUpdateType
	COMPLETE ImportPointsRequestUpdateType
}

func GetImportPointsRequestUpdateTypeEnum() ImportPointsRequestUpdateTypeEnum {
	return ImportPointsRequestUpdateTypeEnum{
		PARTIAL: ImportPointsRequestUpdateType{
			value: "PARTIAL",
		},
		COMPLETE: ImportPointsRequestUpdateType{
			value: "COMPLETE",
		},
	}
}

func (c ImportPointsRequestUpdateType) Value() string {
	return c.value
}

func (c ImportPointsRequestUpdateType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ImportPointsRequestUpdateType) UnmarshalJSON(b []byte) error {
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
