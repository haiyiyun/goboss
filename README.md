#注意事项

1. 海溢云业务运营支撑系统（goboss）使用了mongodb的事务，故必须使用mongodb的replica模式 
1. 使用docker-scripts里的mongo/mongo-replset.sh来运行docker环境的mongodb
1. 使用docker环境运行mongo replset模式时
    要注意在客户端连接环境下的手动添加类似local-mongo-rs-27017等DNS,否则会出现连接失败情况