docker run --name blog-app-db -p 3306:3306 -v "$(pwd)/data:/var/lib/mysql" -e MYSQL_ROOT_PASSWORD=@root -d mysql --character-set-server=utf8mb4 --collation-server=utf8mb4_unicode_ci
