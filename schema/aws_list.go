package schema

import "encoding/json"

type InAWSListRecord struct {
    ContructCommonID string `json:"field_1706862164"`
    ContructID string `json:"field_1706862196"`
    TeikyoID string `json:"field_1706862218"`
    TeikyoName string `json:"field_1706862235"`
    KeiyakuID string `json:"field_1706862251"`
    KeiyakuName string `json:"field_1706862267"`
    ProjectName string `json:"field_1706862364"`
    PayerAccountID string `json:"field_1706862657"`
    AWSAccountID string `json:"field_1706862670"`
    BSCTenat string `json:"field_1706862746"`
    Kind string `json:"field_1706864414"`
    ContructType string `json:"field_1706864444"`
    SupportLevel string `json:"field_1706864480"`
}

type OutAWSList struct {
    RecordId json.Number `json:"recordId" csv:"recordId"`
    ContructID json.Number `json:"ContructID" csv:"ContructID"`
    ContructCommonID json.Number `json:"ContructCommonID" csv:"ContructCommonID"`
    TeikyoID json.Number `json:"TeikyoID" csv:"TeikyoID"`
    TeikyoName string `json:"TeikyoName" csv:"TeikyoName"`
    KeiyakuID json.Number `json:"KeiyakuID" csv:"KeiyakuID"`
    KeiyakuName string `json:"KeiyakuName" csv:"KeiyakuName"`
    ProjectName string `json:"ProjectName" csv:"ProjectName"`

    PayerAccountID json.Number `json:"PayerAccountID" csv:"PayerAccountID"`
    AWSAccountID json.Number `json:"AWSAccountID" csv:"AWSAccountID"`

    BSCTenat string `json:"BSCTenat" csv:"BSCTenat"`
    Kind string `json:"Kind" csv:"Kind"`
    ContructType string `json:"ContructType" csv:"ContructType"`
    SupportLevel string `json:"SupportLevel" csv:"SupportLevel"`
}




