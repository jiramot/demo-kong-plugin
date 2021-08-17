package delivery

import (
    "log"
    "github.com/Kong/go-pdk"
)

type Config struct {
    Apikey string
}

func New() interface{} {
    return &Config{}
}
func (conf Config) Access(kong *pdk.PDK) {
    log.Println("[KeyChecker] Got request")
    key, err := kong.Request.GetQueryArg("key")
    apiKey := conf.Apikey
    if err != nil {
        kong.Log.Err(err.Error())
    }

    respHeaders := make(map[string][]string)
    respHeaders["Content-Type"] = append(respHeaders["Content-Type"], "application/json")

    if apiKey != key {
        kong.Response.Exit(403, "You have no correct key", respHeaders)
        return
    }
}