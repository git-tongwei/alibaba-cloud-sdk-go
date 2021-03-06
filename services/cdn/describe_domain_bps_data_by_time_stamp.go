package cdn

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

func (client *Client) DescribeDomainBpsDataByTimeStamp(request *DescribeDomainBpsDataByTimeStampRequest) (response *DescribeDomainBpsDataByTimeStampResponse, err error) {
	response = CreateDescribeDomainBpsDataByTimeStampResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeDomainBpsDataByTimeStampWithChan(request *DescribeDomainBpsDataByTimeStampRequest) (<-chan *DescribeDomainBpsDataByTimeStampResponse, <-chan error) {
	responseChan := make(chan *DescribeDomainBpsDataByTimeStampResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDomainBpsDataByTimeStamp(request)
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

func (client *Client) DescribeDomainBpsDataByTimeStampWithCallback(request *DescribeDomainBpsDataByTimeStampRequest, callback func(response *DescribeDomainBpsDataByTimeStampResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDomainBpsDataByTimeStampResponse
		var err error
		defer close(result)
		response, err = client.DescribeDomainBpsDataByTimeStamp(request)
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

type DescribeDomainBpsDataByTimeStampRequest struct {
	*requests.RpcRequest
	IspNames      string `position:"Query" name:"IspNames"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	LocationNames string `position:"Query" name:"LocationNames"`
	DomainName    string `position:"Query" name:"DomainName"`
	Action        string `position:"Query" name:"Action"`
	OwnerId       string `position:"Query" name:"OwnerId"`
	AccessKeyId   string `position:"Query" name:"AccessKeyId"`
	TimePoint     string `position:"Query" name:"TimePoint"`
}

type DescribeDomainBpsDataByTimeStampResponse struct {
	*responses.BaseResponse
	RequestId   string `json:"RequestId"`
	DomainName  string `json:"DomainName"`
	TimeStamp   string `json:"TimeStamp"`
	BpsDataList []struct {
		LocationName string `json:"LocationName"`
		IspName      string `json:"IspName"`
		Bps          int64  `json:"Bps"`
	} `json:"BpsDataList"`
}

func CreateDescribeDomainBpsDataByTimeStampRequest() (request *DescribeDomainBpsDataByTimeStampRequest) {
	request = &DescribeDomainBpsDataByTimeStampRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2014-11-11", "DescribeDomainBpsDataByTimeStamp", "", "")
	return
}

func CreateDescribeDomainBpsDataByTimeStampResponse() (response *DescribeDomainBpsDataByTimeStampResponse) {
	response = &DescribeDomainBpsDataByTimeStampResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
