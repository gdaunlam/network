docker stop watchtower;
docker rm watchtower;
docker rmi containrrr/watchtower;
docker stop psclient;
docker rm psclient;
docker rmi packetstream/psclient;
docker run -d --restart=always -e CID=64VT --name psclient packetstream/psclient:latest;
docker run -d --restart=always --name watchtower -v /var/run/docker.sock:/var/run/docker.sock containrrr/watchtower --cleanup --include-stopped --include-restarting --revive-stopped --interval 60 psclient;
