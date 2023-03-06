# mdns-discovery

service config
、、、
docker run --rm \
    --net host \
    --cap-drop ALL \
    --read-only \
    kayuii\discov discov \
        -name service.x1 \
        -service _own_work._tcp \
        -host x1.service.own \
        -port 8080 \
        -ip 192.168.1.1
、、、

client config
、、、
docker run --rm \
    --net host \
    --cap-drop ALL \
    --read-only \
    kayuii\resolv resolv -service _own_work._tcp
、、、
