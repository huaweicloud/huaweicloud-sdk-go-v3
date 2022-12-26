package v3

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/aom/v3/model"
)

type AomClient struct {
	HcClient *http_client.HcHttpClient
}

func NewAomClient(hcClient *http_client.HcHttpClient) *AomClient {
	return &AomClient{HcClient: hcClient}
}

func AomClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

// CreateApp 新增应用
//
// 新增应用。（注：接口目前开放的region为：上海一）。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AomClient) CreateApp(request *model.CreateAppRequest) (*model.CreateAppResponse, error) {
	requestDef := GenReqDefForCreateApp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAppResponse), nil
	}
}

// CreateAppInvoker 新增应用
func (c *AomClient) CreateAppInvoker(request *model.CreateAppRequest) *CreateAppInvoker {
	requestDef := GenReqDefForCreateApp()
	return &CreateAppInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateComponent 新增组件
//
// 新增组件。（注：接口目前开放的region为：上海一）。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AomClient) CreateComponent(request *model.CreateComponentRequest) (*model.CreateComponentResponse, error) {
	requestDef := GenReqDefForCreateComponent()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateComponentResponse), nil
	}
}

// CreateComponentInvoker 新增组件
func (c *AomClient) CreateComponentInvoker(request *model.CreateComponentRequest) *CreateComponentInvoker {
	requestDef := GenReqDefForCreateComponent()
	return &CreateComponentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateEnv 创建环境
//
// 创建环境。（注：接口目前开放的region为：上海一）。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AomClient) CreateEnv(request *model.CreateEnvRequest) (*model.CreateEnvResponse, error) {
	requestDef := GenReqDefForCreateEnv()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateEnvResponse), nil
	}
}

// CreateEnvInvoker 创建环境
func (c *AomClient) CreateEnvInvoker(request *model.CreateEnvRequest) *CreateEnvInvoker {
	requestDef := GenReqDefForCreateEnv()
	return &CreateEnvInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteApp 删除应用
//
// 删除应用。（注：接口目前开放的region为：上海一）。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AomClient) DeleteApp(request *model.DeleteAppRequest) (*model.DeleteAppResponse, error) {
	requestDef := GenReqDefForDeleteApp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAppResponse), nil
	}
}

// DeleteAppInvoker 删除应用
func (c *AomClient) DeleteAppInvoker(request *model.DeleteAppRequest) *DeleteAppInvoker {
	requestDef := GenReqDefForDeleteApp()
	return &DeleteAppInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteComponent 删除组件
//
// 删除组件。（注：接口目前开放的region为：上海一）。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AomClient) DeleteComponent(request *model.DeleteComponentRequest) (*model.DeleteComponentResponse, error) {
	requestDef := GenReqDefForDeleteComponent()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteComponentResponse), nil
	}
}

// DeleteComponentInvoker 删除组件
func (c *AomClient) DeleteComponentInvoker(request *model.DeleteComponentRequest) *DeleteComponentInvoker {
	requestDef := GenReqDefForDeleteComponent()
	return &DeleteComponentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteEnv 删除环境
//
// 删除环境。（注：接口目前开放的region为：上海一）。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AomClient) DeleteEnv(request *model.DeleteEnvRequest) (*model.DeleteEnvResponse, error) {
	requestDef := GenReqDefForDeleteEnv()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteEnvResponse), nil
	}
}

// DeleteEnvInvoker 删除环境
func (c *AomClient) DeleteEnvInvoker(request *model.DeleteEnvRequest) *DeleteEnvInvoker {
	requestDef := GenReqDefForDeleteEnv()
	return &DeleteEnvInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListResourceUnderNode 查询绑定在节点上的资源列表
//
// 查询绑定在节点上的资源列表。（注：接口目前开放的region为：上海一）。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AomClient) ListResourceUnderNode(request *model.ListResourceUnderNodeRequest) (*model.ListResourceUnderNodeResponse, error) {
	requestDef := GenReqDefForListResourceUnderNode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListResourceUnderNodeResponse), nil
	}
}

// ListResourceUnderNodeInvoker 查询绑定在节点上的资源列表
func (c *AomClient) ListResourceUnderNodeInvoker(request *model.ListResourceUnderNodeRequest) *ListResourceUnderNodeInvoker {
	requestDef := GenReqDefForListResourceUnderNode()
	return &ListResourceUnderNodeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowApp 查询应用详情
//
// 获取应用详情。（注：接口目前开放的region为：上海一）。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AomClient) ShowApp(request *model.ShowAppRequest) (*model.ShowAppResponse, error) {
	requestDef := GenReqDefForShowApp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAppResponse), nil
	}
}

// ShowAppInvoker 查询应用详情
func (c *AomClient) ShowAppInvoker(request *model.ShowAppRequest) *ShowAppInvoker {
	requestDef := GenReqDefForShowApp()
	return &ShowAppInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAppByName 查询应用详情
//
// 获取应用详情。（注：接口目前开放的region为：上海一）。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AomClient) ShowAppByName(request *model.ShowAppByNameRequest) (*model.ShowAppByNameResponse, error) {
	requestDef := GenReqDefForShowAppByName()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAppByNameResponse), nil
	}
}

// ShowAppByNameInvoker 查询应用详情
func (c *AomClient) ShowAppByNameInvoker(request *model.ShowAppByNameRequest) *ShowAppByNameInvoker {
	requestDef := GenReqDefForShowAppByName()
	return &ShowAppByNameInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowComponent 查询组件详情
//
// 查询组件详情。（注：接口目前开放的region为：上海一）。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AomClient) ShowComponent(request *model.ShowComponentRequest) (*model.ShowComponentResponse, error) {
	requestDef := GenReqDefForShowComponent()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowComponentResponse), nil
	}
}

// ShowComponentInvoker 查询组件详情
func (c *AomClient) ShowComponentInvoker(request *model.ShowComponentRequest) *ShowComponentInvoker {
	requestDef := GenReqDefForShowComponent()
	return &ShowComponentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowComponentByName 查询组件详情
//
// 查询组件详情。（注：接口目前开放的region为：上海一）。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AomClient) ShowComponentByName(request *model.ShowComponentByNameRequest) (*model.ShowComponentByNameResponse, error) {
	requestDef := GenReqDefForShowComponentByName()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowComponentByNameResponse), nil
	}
}

// ShowComponentByNameInvoker 查询组件详情
func (c *AomClient) ShowComponentByNameInvoker(request *model.ShowComponentByNameRequest) *ShowComponentByNameInvoker {
	requestDef := GenReqDefForShowComponentByName()
	return &ShowComponentByNameInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowEnv 查询环境详情
//
// 查询环境详情。（注：接口目前开放的region为：上海一）。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AomClient) ShowEnv(request *model.ShowEnvRequest) (*model.ShowEnvResponse, error) {
	requestDef := GenReqDefForShowEnv()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowEnvResponse), nil
	}
}

// ShowEnvInvoker 查询环境详情
func (c *AomClient) ShowEnvInvoker(request *model.ShowEnvRequest) *ShowEnvInvoker {
	requestDef := GenReqDefForShowEnv()
	return &ShowEnvInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowEnvByName 查询环境详情
//
// 查询环境详情。（注：接口目前开放的region为：上海一）。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AomClient) ShowEnvByName(request *model.ShowEnvByNameRequest) (*model.ShowEnvByNameResponse, error) {
	requestDef := GenReqDefForShowEnvByName()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowEnvByNameResponse), nil
	}
}

// ShowEnvByNameInvoker 查询环境详情
func (c *AomClient) ShowEnvByNameInvoker(request *model.ShowEnvByNameRequest) *ShowEnvByNameInvoker {
	requestDef := GenReqDefForShowEnvByName()
	return &ShowEnvByNameInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateApp 修改应用
//
// 修改应用。（注：接口目前开放的region为：上海一）。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AomClient) UpdateApp(request *model.UpdateAppRequest) (*model.UpdateAppResponse, error) {
	requestDef := GenReqDefForUpdateApp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAppResponse), nil
	}
}

// UpdateAppInvoker 修改应用
func (c *AomClient) UpdateAppInvoker(request *model.UpdateAppRequest) *UpdateAppInvoker {
	requestDef := GenReqDefForUpdateApp()
	return &UpdateAppInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateComponent 修改组件
//
// 修改组件。（注：接口目前开放的region为：上海一）。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AomClient) UpdateComponent(request *model.UpdateComponentRequest) (*model.UpdateComponentResponse, error) {
	requestDef := GenReqDefForUpdateComponent()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateComponentResponse), nil
	}
}

// UpdateComponentInvoker 修改组件
func (c *AomClient) UpdateComponentInvoker(request *model.UpdateComponentRequest) *UpdateComponentInvoker {
	requestDef := GenReqDefForUpdateComponent()
	return &UpdateComponentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateEnv 修改环境
//
// 修改环境。（注：接口目前开放的region为：上海一）。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AomClient) UpdateEnv(request *model.UpdateEnvRequest) (*model.UpdateEnvResponse, error) {
	requestDef := GenReqDefForUpdateEnv()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateEnvResponse), nil
	}
}

// UpdateEnvInvoker 修改环境
func (c *AomClient) UpdateEnvInvoker(request *model.UpdateEnvRequest) *UpdateEnvInvoker {
	requestDef := GenReqDefForUpdateEnv()
	return &UpdateEnvInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
