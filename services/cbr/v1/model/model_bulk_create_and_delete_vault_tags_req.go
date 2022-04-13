package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

//
type BulkCreateAndDeleteVaultTagsReq struct {
	// 标签列表。  tags不允许为空列表。  tags中最多包含10个key。  tags中key不允许重复。

	Tags *[]Tag `json:"tags,omitempty"`
	// 系统标签列表。  op_service权限可以访问，和tags二选一。  目前TMS调用时只包含一个resource_tag结构体 ，key固定为：_sys_enterprise_project_id。  value是UUID或0,value为0表示默认企业项目。  现在仅支持create操作。

	SysTags *[]SysTag `json:"sys_tags,omitempty"`
	// 操作标识：仅限于create（创建）、delete（删除）

	Action BulkCreateAndDeleteVaultTagsReqAction `json:"action"`
}

func (o BulkCreateAndDeleteVaultTagsReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BulkCreateAndDeleteVaultTagsReq struct{}"
	}

	return strings.Join([]string{"BulkCreateAndDeleteVaultTagsReq", string(data)}, " ")
}

type BulkCreateAndDeleteVaultTagsReqAction struct {
	value string
}

type BulkCreateAndDeleteVaultTagsReqActionEnum struct {
	CREATE BulkCreateAndDeleteVaultTagsReqAction
	DELETE BulkCreateAndDeleteVaultTagsReqAction
}

func GetBulkCreateAndDeleteVaultTagsReqActionEnum() BulkCreateAndDeleteVaultTagsReqActionEnum {
	return BulkCreateAndDeleteVaultTagsReqActionEnum{
		CREATE: BulkCreateAndDeleteVaultTagsReqAction{
			value: "create",
		},
		DELETE: BulkCreateAndDeleteVaultTagsReqAction{
			value: " delete",
		},
	}
}

func (c BulkCreateAndDeleteVaultTagsReqAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BulkCreateAndDeleteVaultTagsReqAction) UnmarshalJSON(b []byte) error {
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
