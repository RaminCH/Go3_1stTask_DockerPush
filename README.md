# Go3_1stTask_DockerPush

sudo docker login
sudo docker build -t raminch/go_api_items . # -> here . mean we take all data from current folder
sudo docker images # should be image named --> raminch/go_api_items
sudo docker push raminch/go_api_items:latest

# sudo docker run -p 8000:8080 -it tester # 'ports' -> left localhost / right docker /// 'tester' is the image name
