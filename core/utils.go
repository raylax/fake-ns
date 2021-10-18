package core

import "strings"

func GetPrefix(domain string, domainSuffix string) string {
    if strings.HasSuffix(domain, domainSuffix) {
        length := len(domain) - len(domainSuffix)
        if length > 1 {
            return domain[:length-1]
        }
    }
    return ""
}
