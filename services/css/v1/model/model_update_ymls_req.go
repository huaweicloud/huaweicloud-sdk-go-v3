package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type UpdateYmlsReq struct {
	Edit *UpdateYmlsReqEdit `json:"edit"`

	// 节点类型 目前koosearch集群涉及不同类型的节点。 kos: koosearch的搜索中控节点 kos-doc: koosearch的文档解析节点
	InstType *UpdateYmlsReqInstType `json:"inst_type,omitempty"`
}

func (o UpdateYmlsReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateYmlsReq struct{}"
	}

	return strings.Join([]string{"UpdateYmlsReq", string(data)}, " ")
}

type UpdateYmlsReqInstType struct {
	value string
}

type UpdateYmlsReqInstTypeEnum struct {
	KOS     UpdateYmlsReqInstType
	KOS_DOC UpdateYmlsReqInstType
}

func GetUpdateYmlsReqInstTypeEnum() UpdateYmlsReqInstTypeEnum {
	return UpdateYmlsReqInstTypeEnum{
		KOS: UpdateYmlsReqInstType{
			value: "kos",
		},
		KOS_DOC: UpdateYmlsReqInstType{
			value: "kos-doc",
		},
	}
}

func (c UpdateYmlsReqInstType) Value() string {
	return c.value
}

func (c UpdateYmlsReqInstType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateYmlsReqInstType) UnmarshalJSON(b []byte) error {
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
