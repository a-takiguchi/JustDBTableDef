package physi_logi_def

var tables = map[string]string{
    "domain-ssl":  "table_1704780398",
}

func GetLogiName(alias string) string {
    return tables[alias]
}
