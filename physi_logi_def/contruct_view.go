package physi_logi_def
import "encoding/json"

type OutContructView struct {
    ContructID json.Number `json:"ContructID"`
    KeiyakuID json.Number `json:"KeiyakuID"`
    KeiyakuName string `json:"KeiyakuName"`
    KeiyakuZACCode json.Number `json:"KeiyakuZACCode"`
    TeikyoID json.Number `json:"TeikyoID"`
    TeikyoName string `json:"TeikyoName"`
    TeikyoZACCode json.Number `json:"TeikyoZACCode"`
    ProjectName string `json:"ProjectName"`
    StartDate string `json:"StartDate"`
    EndDate string `json:"EndDate"`
    KeiyakuShoNumber string `json:"KeiyakuShoNumber"`
    ContructCommonID json.Number `json:"ContructCommonID"`
    ShohinID string `json:"ShohinID"`
    ShohinName string `json:"ShohinName"`
    PlanID string `json:"PlanID"`
    PlanName string `json:"PlanName"`
}

const ContructID = "ContructID"
const KeiyakuID = "KeiyakuID"
const KeiyakuName = "KeiyakuName"
const KeiyakuZACCode = "KeiyakuZACCode"
const TeikyoID = "TeikyoID"
const TeikyoName = "TeikyoName"
const TeikyoZACCode = "TeikyoZACCode"
const ProjectName = "ProjectName"
const StartDate = "StartDate"
const EndDate = "EndDate"
const KeiyakuShoNumber = "KeiyakuShoNumber"
const ContructCommonID = "ContructCommonID"
const ShohinID = "ShohinID"
const ShohinName = "ShohinName"
const PlanID = "PlanID"
const PlanName = "PlanName"

func GetContructViewHeader() []string {
    return []string{
        ContructID,
        KeiyakuID,
        KeiyakuName,
        KeiyakuZACCode,
        TeikyoID,
        TeikyoName,
        TeikyoZACCode,
        ProjectName,
        StartDate,
        EndDate,
        KeiyakuShoNumber,
        ContructCommonID,
        ShohinID,
        ShohinName,
        PlanID,
        PlanName,
    }
}
