build with custom module
```
xcaddy build --with github.com/niole/test-caddy-module=.
```

run
```
docker run --rm -p 8080:80 caddy-test:1
```

test
```
curl localhost:8080 # should print out "NIOLE" in the caddy container logs
```

dev
```
xcaddy run
```
