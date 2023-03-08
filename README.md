# mdns-discovery

service config
```
#!/bin/bash
ipaddr=$(ifconfig enp3s0 | grep 'inet ' | awk '{print $2}')
docker run --rm \
    --net host \
    --cap-drop ALL \
    --read-only \
    kayuii/mdnscli service \
        --name service.x1 \
        --service _own_work._tcp \
        --host x1.service.own \
        --port 8080 \
        --ip $ipaddr
```

client config
```
docker run --rm \
    --net host \
    --cap-drop ALL \
    --read-only \
    kayuii/mdnscli client --service _own_work._tcp
```
