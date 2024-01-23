package json_schema

import (
    "fmt"
    "domain_ssl"
    "contruct_view"
)


func GetJsonSchema(name string) interface{} {
    switch name {
    case "domain_ssl":
        return domain_ssl.DomainSSL{}
    case "contruct-view":
        return contruct_view.ContructView{}

    default:
        panic(fmt.Sprintf("cannot find: %s json schema", name))
    }
}
