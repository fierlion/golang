docker run --name mysqlgo -e MYSQL_ROOT_PASSWORD=mysqlgo \
                          -e MYSQL_DATABASE=mysqlgotest \
                          -e MYSQL_USER=mysqlgo \
                          -e MYSQL_PASSWORD=mysqlgopw \
                          -d mysql
