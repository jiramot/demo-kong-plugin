# Kong
## Install OSX
```
brew tap kong/kong
brew install kong
```
## Creating a Declarative Configuration File
```
kong config -c kong.conf init
```

## Custom API Key
https://konghq.com/blog/kong-gateway-go-plugin

```
docker build -t kong-demo .
```
```
docker run -ti --rm --name kong-go-plugins \
-e "KONG_DATABASE=off" \
-e "KONG_DECLARATIVE_CONFIG=/tmp/config.yml" \
-e "KONG_PLUGINS=bundled,key-checker" \
-e "KONG_PLUGINSERVER_NAMES=key-checker" \
-e "KONG_PLUGINSERVER_KEY_CHECKER_START_CMD=/usr/local/bin/key-checker" \
-e "KONG_PLUGINSERVER_KEY_CHECKER_QUERY_CMD=/usr/local/bin/key-checker -dump" \
-e "KONG_PROXY_LISTEN=0.0.0.0:8000" \
-e "KONG_LOG_LEVEL=debug" \
-p 8000:8000 \
 kong-demo
```

## Run
```
curl --location --request GET '0.0.0.0:8000?key=invalidconsumerkey'
curl --location --request GET '0.0.0.0:8000?key=mysecretconsumerkey'

```