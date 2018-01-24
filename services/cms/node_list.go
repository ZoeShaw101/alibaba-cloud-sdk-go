package cms

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

func (client *Client) NodeList(request *NodeListRequest) (response *NodeListResponse, err error) {
	response = CreateNodeListResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) NodeListWithChan(request *NodeListRequest) (<-chan *NodeListResponse, <-chan error) {
	responseChan := make(chan *NodeListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.NodeList(request)
		responseChan <- response
		errChan <- err
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

func (client *Client) NodeListWithCallback(request *NodeListRequest, callback func(response *NodeListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *NodeListResponse
		var err error
		defer close(result)
		response, err = client.NodeList(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

type NodeListRequest struct {
	*requests.RpcRequest
	HostName      string           `position:"Query" name:"HostName"`
	InstanceIds   string           `position:"Query" name:"InstanceIds"`
	PageSize      requests.Integer `position:"Query" name:"PageSize"`
	KeyWord       string           `position:"Query" name:"KeyWord"`
	UserId        requests.Integer `position:"Query" name:"UserId"`
	SerialNumbers string           `position:"Query" name:"SerialNumbers"`
	PageNumber    requests.Integer `position:"Query" name:"PageNumber"`
	Status        string           `position:"Query" name:"Status"`
}

type NodeListResponse struct {
	*responses.BaseResponse
	ErrorCode    int    `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	Success      bool   `json:"Success" xml:"Success"`
	RequestId    string `json:"RequestId" xml:"RequestId"`
	PageNumber   int    `json:"PageNumber" xml:"PageNumber"`
	PageSize     int    `json:"PageSize" xml:"PageSize"`
	PageTotal    int    `json:"PageTotal" xml:"PageTotal"`
	Total        int    `json:"Total" xml:"Total"`
	Nodes        struct {
		Node []struct {
			InstanceId       string `json:"InstanceId" xml:"InstanceId"`
			SerialNumber     string `json:"SerialNumber" xml:"SerialNumber"`
			HostName         string `json:"HostName" xml:"HostName"`
			AliUid           int    `json:"AliUid" xml:"AliUid"`
			OperatingSystem  string `json:"OperatingSystem" xml:"OperatingSystem"`
			IpGroup          string `json:"IpGroup" xml:"IpGroup"`
			Region           string `json:"Region" xml:"Region"`
			TianjimonVersion string `json:"TianjimonVersion" xml:"TianjimonVersion"`
			EipAddress       string `json:"EipAddress" xml:"EipAddress"`
			EipId            string `json:"EipId" xml:"EipId"`
			AliyunHost       bool   `json:"AliyunHost" xml:"AliyunHost"`
			NatIp            string `json:"NatIp" xml:"NatIp"`
			NetworkType      string `json:"NetworkType" xml:"NetworkType"`
		} `json:"Node" xml:"Node"`
	} `json:"Nodes" xml:"Nodes"`
}

func CreateNodeListRequest() (request *NodeListRequest) {
	request = &NodeListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2017-03-01", "NodeList", "cms", "openAPI")
	return
}

func CreateNodeListResponse() (response *NodeListResponse) {
	response = &NodeListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
