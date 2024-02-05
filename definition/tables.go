package definition

const DomainSSL string = "domain_ssl"
const ContructView string = "contruct_view"
const SlackUser string = "slack_user"
const AWSList string = "aws_list"


type AccessInfo struct {
    tableID string
    filterID string
    panelID string
}

var tables = map[string]AccessInfo{
    DomainSSL: {
        tableID: "table_1704780398",
        panelID:  "panel_1704780386",
        filterID: "filter_1704786328",
    },
    ContructView: {
        tableID: "table_1705998836",
        panelID:  "panel_1705998821",
        filterID: "filter_1706000468",
    },
    SlackUser: {
        tableID: "table_1705543400",
        panelID:  "panel_1705543387",
        filterID: "filter_1706249642",
    },

    AWSList: {
        tableID: "table_1706862126",
        panelID:  "panel_1706862115",
        filterID: "",
    },
}

func GetTableID(alias string) string {
    return tables[alias].tableID
}
func GetPanelID(alias string) string {
    return tables[alias].panelID
}
func GetFilterID(alias string) string {
    return tables[alias].filterID
}
