package rds

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

// Template is a nested struct in rds response
type Template struct {
	TotalConsume    int     `json:"TotalConsume" xml:"TotalConsume"`
	TotalUpdateRows int     `json:"TotalUpdateRows" xml:"TotalUpdateRows"`
	SqlType         string  `json:"SqlType" xml:"SqlType"`
	Template        string  `json:"Template" xml:"Template"`
	TotalCounts     int     `json:"TotalCounts" xml:"TotalCounts"`
	TotalScanRows   int     `json:"TotalScanRows" xml:"TotalScanRows"`
	AvgUpdateRows   float64 `json:"AvgUpdateRows" xml:"AvgUpdateRows"`
	TemplateHash    string  `json:"TemplateHash" xml:"TemplateHash"`
	AvgConsume      float64 `json:"AvgConsume" xml:"AvgConsume"`
	AvgScanRows     float64 `json:"AvgScanRows" xml:"AvgScanRows"`
}
