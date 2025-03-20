package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ImageCacheDetail 镜像缓存信息。
type ImageCacheDetail struct {

	// 镜像缓存名称。
	Name string `json:"name"`

	// 镜像缓存ID。
	Id string `json:"id"`

	// 镜像缓存创建时间戳。
	CreatedAt string `json:"created_at"`

	// 镜像缓存中的容器镜像列表。
	Images []string `json:"images"`

	// 镜像缓存后端对应的存储盘大小，单位GiB。
	ImageCacheSize int32 `json:"image_cache_size"`

	// **参数解释：** 镜像缓存有效的天数,不设置或值为0表示永久有效。镜像缓存到达有效期后自动过期并删除。 **约束限制：** 不涉及 **取值范围：** [0-10000] **默认取值：** 0
	RetentionDays int32 `json:"retention_days"`

	BuildingConfig *ImageCacheBuildingConfig `json:"building_config"`

	// **参数解释：** 镜像缓存的状态。 **约束限制：** 不涉及 **取值范围：** - Available： 可用状态，表示当前镜像缓存正常可用。 - Unavailable：不可用，表示镜像缓存状态异常或过期，不可使用。 - Creating：创建中，表示镜像缓存正在创建中。 - Deleting：删除中，表示镜像缓存正在删除中。 - Failed：创建失败，表示镜像缓存创建失败。 **默认取值：** 不涉及
	Status ImageCacheDetailStatus `json:"status"`

	// 镜像缓存创建失败或异常的错误信息。
	Message *string `json:"message,omitempty"`
}

func (o ImageCacheDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageCacheDetail struct{}"
	}

	return strings.Join([]string{"ImageCacheDetail", string(data)}, " ")
}

type ImageCacheDetailStatus struct {
	value string
}

type ImageCacheDetailStatusEnum struct {
	AVAILABLE   ImageCacheDetailStatus
	UNAVAILABLE ImageCacheDetailStatus
	CREATING    ImageCacheDetailStatus
	DELETING    ImageCacheDetailStatus
	FAILED      ImageCacheDetailStatus
}

func GetImageCacheDetailStatusEnum() ImageCacheDetailStatusEnum {
	return ImageCacheDetailStatusEnum{
		AVAILABLE: ImageCacheDetailStatus{
			value: "Available",
		},
		UNAVAILABLE: ImageCacheDetailStatus{
			value: "Unavailable",
		},
		CREATING: ImageCacheDetailStatus{
			value: "Creating",
		},
		DELETING: ImageCacheDetailStatus{
			value: "Deleting",
		},
		FAILED: ImageCacheDetailStatus{
			value: "Failed",
		},
	}
}

func (c ImageCacheDetailStatus) Value() string {
	return c.value
}

func (c ImageCacheDetailStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ImageCacheDetailStatus) UnmarshalJSON(b []byte) error {
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
