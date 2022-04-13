package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// Request Object
type DeleteImageSyncRepoRequest struct {
	// 消息体的类型（格式），下方类型可任选其一使用： application/json;charset=utf-8 application/json

	ContentType DeleteImageSyncRepoRequestContentType `json:"Content-Type"`
	// 组织名称。小写字母开头，后面跟小写字母、数字、小数点、下划线或中划线（其中下划线最多允许连续两个，小数点、下划线、中划线不能直接相连），小写字母或数字结尾，1-64个字符。

	Namespace string `json:"namespace"`
	// 镜像仓库名称

	Repository string `json:"repository"`

	Body *DeleteImageSyncRepoRequestBody `json:"body,omitempty"`
}

func (o DeleteImageSyncRepoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteImageSyncRepoRequest struct{}"
	}

	return strings.Join([]string{"DeleteImageSyncRepoRequest", string(data)}, " ")
}

type DeleteImageSyncRepoRequestContentType struct {
	value string
}

type DeleteImageSyncRepoRequestContentTypeEnum struct {
	APPLICATION_JSONCHARSETUTF_8 DeleteImageSyncRepoRequestContentType
	APPLICATION_JSON             DeleteImageSyncRepoRequestContentType
}

func GetDeleteImageSyncRepoRequestContentTypeEnum() DeleteImageSyncRepoRequestContentTypeEnum {
	return DeleteImageSyncRepoRequestContentTypeEnum{
		APPLICATION_JSONCHARSETUTF_8: DeleteImageSyncRepoRequestContentType{
			value: "application/json;charset=utf-8",
		},
		APPLICATION_JSON: DeleteImageSyncRepoRequestContentType{
			value: "application/json",
		},
	}
}

func (c DeleteImageSyncRepoRequestContentType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteImageSyncRepoRequestContentType) UnmarshalJSON(b []byte) error {
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
