package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowSyncJobRequest Request Object
type ShowSyncJobRequest struct {

	// 消息体的类型（格式），下方类型可任选其一使用： application/json;charset=utf-8 application/json
	ContentType ShowSyncJobRequestContentType `json:"Content-Type"`

	// 组织名称。小写字母开头，后面跟小写字母、数字、小数点、下划线或中划线（其中下划线最多允许连续两个，小数点、下划线、中划线不能直接相连），小写字母或数字结尾，1-64个字符。
	Namespace string `json:"namespace"`

	// 镜像仓库名称
	Repository string `json:"repository"`

	// 应填写 limit::{limit}|offset::{offset}|order::{order} ,其中{limit}为返回条数,{offset}为起始索引,注意：offset和limit参数需要配套使用，且必选。{order}为排序类型（可选），可设置为desc（降序）、asc（升序）
	Filter string `json:"filter"`
}

func (o ShowSyncJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSyncJobRequest struct{}"
	}

	return strings.Join([]string{"ShowSyncJobRequest", string(data)}, " ")
}

type ShowSyncJobRequestContentType struct {
	value string
}

type ShowSyncJobRequestContentTypeEnum struct {
	APPLICATION_JSONCHARSETUTF_8 ShowSyncJobRequestContentType
	APPLICATION_JSON             ShowSyncJobRequestContentType
}

func GetShowSyncJobRequestContentTypeEnum() ShowSyncJobRequestContentTypeEnum {
	return ShowSyncJobRequestContentTypeEnum{
		APPLICATION_JSONCHARSETUTF_8: ShowSyncJobRequestContentType{
			value: "application/json;charset=utf-8",
		},
		APPLICATION_JSON: ShowSyncJobRequestContentType{
			value: "application/json",
		},
	}
}

func (c ShowSyncJobRequestContentType) Value() string {
	return c.value
}

func (c ShowSyncJobRequestContentType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowSyncJobRequestContentType) UnmarshalJSON(b []byte) error {
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
