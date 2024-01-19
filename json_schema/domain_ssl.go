/*
[
{"tableName":"table_1704780398",
"recordId":2,
"reactions":[],
"comments":[],
"createdBy":["atsuhiro.takiguchi","滝口 敦弘","atsuhiro.takiguchi@beex-inc.com","マネージドサービス部","メンバー"],
"createdAt":"2024-01-09T16:20:17.380+09:00",
"updatedBy":["atsuhiro.takiguchi","滝口 敦弘","atsuhiro.takiguchi@beex-inc.com","マネージドサービス部","メンバー"],
"updatedAt":"2024-01-09T16:20:17.380+09:00",
"workflowName":"",
"insertUser":[],
"substituteUser":[],
"nextApprovalUsers":[],
"submittedAt":"",
"lastApprovalUser":[],
"approveAt":[""],
"approvalUser":[],
"approvalStatus":"",
"recordLink":"https://beex.just-db.com/sites/view?table=10338&record=2",

	"record":{
	    "field_1704780408":"904551724887",
	    "field_1704781328":"お名前.com(ncp)",
	    "field_1704783209":{"days":168,"hours":9,"minutes":17,"wholeSeconds":0,"microSeconds":0},
	    "field_1704780847":"www.tpocare.com",
	    "field_1704780864":"2024-07-03T00:00+09:00",
	    "field_1704780832":"",
	    "field_1704781767":[],
	    "field_1704781365":["0","対応不要（運用中）"],
	    "field_1704781345":"EC2(nginx)",
	    "field_1704782697":["","2","","2"]}
	}

]
*/
package json_schema

import (
	"encoding/json"
)

type DomainSSL struct {
    TableName string `json:"tableName"`
    RecordId json.Number `json:"recordId"`
    Reactions []string `json:"reactions"`
    Comments []string `json:"comments"`
    CreatedBy []string `json:"createdBy"`
    CreatedAt string `json:"createdAt"`
    UpdatedBy []string `json:"updatedBy"`
    UpdatedAt string `json:"updatedAt"`
    WorkflowName string `json:"workflowName"`
    InsertUser []string `json:"insertUser"`
    SubstituteUser []string `json:"substituteUser"`
    NextApprovalUsers []string `json:"nextApprovalUsers"`
    SubmittedAt string `json:"submittedAt"`
    LastApprovalUser []string `json:"lastApprovalUser"`
    ApproveAt []string `json:"approveAt"`
    ApprovalUser []string `json:"approvalUser"`
    ApprovalStatus string `json:"approvalStatus"`
    RecordLink string `json:"recordLink"`
    Record Record `json:"record"`
}

type Record struct {
    Project string `json:"field_1684834053_K"`
    Customer string `json:"field_1670909964_K22"`
    AWSAccountID string `json:"field_1704780408"`
    Registrar string `json:"field_1704781328"`
    RemainDays Day `json:"field_1704783209"`
    CommonName string `json:"field_1704780847"`
    ExpireDate string `json:"field_1704780864"`
    DomainName string `json:"field_1704780832"`
    PIC []string `json:"field_1704781767"`
    Status []string `json:"field_1704781365"`
    Note string `json:"field_1704781345"`
    ID []string `json:"field_1704782697"`
}

type Day struct {
    Days json.Number `json:"days"`
    Hours json.Number `json:"hours"`
    Minutes json.Number `json:"minutes"`
    WholeSeconds json.Number `json:"wholeSeconds"`
    MicorSeconds json.Number `json:"micorSeconds"`
}

