docker run -e "MYSQL_ROOT_PASSWORD=GwVE7uc2eYT73FNG" --name goapi-sql-dev -p 3307:3306 -v goapi-volume-dev:/var/lib/mysql -d --net goapi-network-dev mysql:latest