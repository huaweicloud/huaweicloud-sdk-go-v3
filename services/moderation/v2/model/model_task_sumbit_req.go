package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

//
type TaskSumbitReq struct {
	// 图片的URL路径，目前支持： - 公网HTTP/HTTPS URL - 华为云OBS提供的URL，使用OBS数据需要进行授权。包括对服务授权、临时授权、匿名公开授权。详请参见[配置OBS访问权限](https://support.huaweicloud.com/api-moderation/moderation_03_0020.html)。 > 图片的URL路径列表最多支持500个URL地址。接口响应时间依赖图片的下载时间，如果图片下载时间过长，会返回接口调用失败。请保证被检测图片所在的存储服务稳定可靠，建议您使用华为云OBS存储。

	Urls []string `json:"urls"`
	// 检测场景。  - politics：是否涉及政治人物的检测。  - terrorism：是否包含涉政敏感人物、涉政暴恐元素的检测。  - porn：是否包含涉黄内容元素的检测。  可通过配置上述场景，来完成对应场景元素的检测。  为空或无此参数时默认检测politics和terrorism(不包含porn)。

	Categories *[]TaskSumbitReqCategories `json:"categories,omitempty"`
}

func (o TaskSumbitReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskSumbitReq struct{}"
	}

	return strings.Join([]string{"TaskSumbitReq", string(data)}, " ")
}

type TaskSumbitReqCategories struct {
	value string
}

type TaskSumbitReqCategoriesEnum struct {
	POLITICS  TaskSumbitReqCategories
	PORN      TaskSumbitReqCategories
	TERRORISM TaskSumbitReqCategories
}

func GetTaskSumbitReqCategoriesEnum() TaskSumbitReqCategoriesEnum {
	return TaskSumbitReqCategoriesEnum{
		POLITICS: TaskSumbitReqCategories{
			value: "politics",
		},
		PORN: TaskSumbitReqCategories{
			value: "porn",
		},
		TERRORISM: TaskSumbitReqCategories{
			value: "terrorism",
		},
	}
}

func (c TaskSumbitReqCategories) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TaskSumbitReqCategories) UnmarshalJSON(b []byte) error {
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
