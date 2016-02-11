# pull images
docker pull mongo
docker pull nsqio/nsq

# activate mongodb
docker run --name mongod -u mongodb -p 27017:27017 mongo sh -c 'mongod'

# actiate nsq
docker run --name lookupd -p 4160:4160 -p 4161:4161 nsqio/nsq /nsqlookupd
docker run --name nsqd -p 4150:4150 -p 4151:4151 nsqio/nsq /nsqd --broadcast-address=172.17.0.1 --lookupd-tcp-address=172.17.0.1:4160
