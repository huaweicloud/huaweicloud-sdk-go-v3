package model

import (
	"encoding/json"

	"strings"
)

// 添加设备结构体。
type AddDevice struct {
	// 设备ID，用于唯一标识一个设备。如果携带该参数，平台将设备ID设置为该参数值；如果不携带该参数，设备ID由物联网平台分配获得，生成规则为\"product_id\" + \"_\" + \"node_id\"拼接而成。

	DeviceId *string `json:"device_id,omitempty"`
	// 设备标识码，通常使用IMEI、MAC地址或Serial No作为node_id。（注意:NB设备由于模组烧录信息后无法配置，所以NB设备会校验node_id全局唯一。）

	NodeId string `json:"node_id"`
	// 设备名称。

	DeviceName *string `json:"device_name,omitempty"`
	// 设备关联的产品ID，用于唯一标识一个产品模型，在管理门户导入产品模型后由平台分配获得。

	ProductId string `json:"product_id"`

	AuthInfo *AuthInfo `json:"auth_info,omitempty"`
	// 设备的描述信息。

	Description *string `json:"description,omitempty"`
	// 网关ID，用于标识设备所属的父设备，即父设备的设备ID。携带该参数时，表示在该父设备下创建一个子设备，这个子设备不与平台直连，此时必须保证这个父设备在平台已存在，创建成功后子设备的gateway_id等于该参数值；不携带该参数时，表示创建一个和平台直连的设备，创建成功后设备的device_id和gateway_id一致。注意：当前平台最多支持二级子设备。

	GatewayId *string `json:"gateway_id,omitempty"`
	// 资源空间ID。此参数为非必选参数，存在多资源空间的用户需要使用该接口时，建议携带该参数指定创建的设备归属到哪个资源空间下，否则创建的设备将会归属到[默认资源空间](https://support.huaweicloud.com/usermanual-iothub/iot_01_0006.html#section0)下。

	AppId *string `json:"app_id,omitempty"`
	// 设备扩展信息。用户可以自定义任何想要的扩展信息，如果在创建设备时为子设备指定该字段，将会通过MQTT接口“平台通知网关子设备新增“将该信息通知给网关。字段值大小上限为1K。

	ExtensionInfo *interface{} `json:"extension_info,omitempty"`
	// 设备初始配置。用户使用该字段可以为设备指定初始配置，指定后将会根据service_id和desired设置的属性值与产品中对应属性的默认值比对，如果不同，则将以shadow字段中设置的属性值为准写入到设备影子中。service_id的值和desired内的属性必须是profile中定义的。

	Shadow *[]InitialDesired `json:"shadow,omitempty"`
}

func (o AddDevice) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "AddDevice struct{}"
	}

	return strings.Join([]string{"AddDevice", string(data)}, " ")
}
