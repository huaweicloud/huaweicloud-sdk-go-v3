package v1

import (
	http_client "github.com/RandolphCYG/huaweicloud-sdk-go-v3/core"

	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/services/imagesearch/v1/model"
)

type ImageSearchClient struct {
	HcClient *http_client.HcHttpClient
}

func NewImageSearchClient(hcClient *http_client.HcHttpClient) *ImageSearchClient {
	return &ImageSearchClient{HcClient: hcClient}
}

func ImageSearchClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

//添加图片 将图片导入图片索引库，该图片可以是同一区域OBS桶内的图片或请求消息体里的图片，只有导入图片索引库的图片方可被搜索到。  > - 添加或搜索的图片存储在OBS的桶中时，需要对OBS的桶授权。在图像搜索服务管理控制台“实例管理”页面，单击实例操作列的“离线导入”，进入“离线导入”页面。选择存放数据的OBS桶，单击“授权”按钮，字体显示为灰色即完成OBS授权访问。
func (c *ImageSearchClient) RunAddPicture(request *model.RunAddPictureRequest) (*model.RunAddPictureResponse, error) {
	requestDef := GenReqDefForRunAddPicture()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunAddPictureResponse), nil
	}
}

//通过图片路径查询索引库中对应图片是否存在。
func (c *ImageSearchClient) RunCheckPicture(request *model.RunCheckPictureRequest) (*model.RunCheckPictureResponse, error) {
	requestDef := GenReqDefForRunCheckPicture()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunCheckPictureResponse), nil
	}
}

//创建实例，实例中会生成图片索引库，用来存放图片特征。
func (c *ImageSearchClient) RunCreateInstance(request *model.RunCreateInstanceRequest) (*model.RunCreateInstanceResponse, error) {
	requestDef := GenReqDefForRunCreateInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunCreateInstanceResponse), nil
	}
}

//删除已存在的实例。
func (c *ImageSearchClient) RunDeleteInstance(request *model.RunDeleteInstanceRequest) (*model.RunDeleteInstanceResponse, error) {
	requestDef := GenReqDefForRunDeleteInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunDeleteInstanceResponse), nil
	}
}

//通过图片路径删除索引库中对应图片。
func (c *ImageSearchClient) RunDeletePicture(request *model.RunDeletePictureRequest) (*model.RunDeletePictureResponse, error) {
	requestDef := GenReqDefForRunDeletePicture()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunDeletePictureResponse), nil
	}
}

//修改图像索引库中已存在的图片信息。
func (c *ImageSearchClient) RunModifyPicture(request *model.RunModifyPictureRequest) (*model.RunModifyPictureResponse, error) {
	requestDef := GenReqDefForRunModifyPicture()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunModifyPictureResponse), nil
	}
}

//查看用户指定实例详细信息。
func (c *ImageSearchClient) RunQueryInstance(request *model.RunQueryInstanceRequest) (*model.RunQueryInstanceResponse, error) {
	requestDef := GenReqDefForRunQueryInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunQueryInstanceResponse), nil
	}
}

//从图片索引库中搜索相似图片。
func (c *ImageSearchClient) RunSearchPicture(request *model.RunSearchPictureRequest) (*model.RunSearchPictureResponse, error) {
	requestDef := GenReqDefForRunSearchPicture()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunSearchPictureResponse), nil
	}
}
