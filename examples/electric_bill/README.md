# 北邮电费查询

北邮电费查询是一个CLI工具，配合crontab等组件可以很方便的实现每日电量查询。在本实例中使用[server酱 Turbo](https://sct.ftqq.com/)向终端分发消息。

## 开始

在example/electric_bill目录下使用如下命令构建可执行程序

```shell
go build
```

构建结束后会得到一个名为electric_bill的可执行程序，执行如下命令查询电量

```shell
./electric_bill -u={USERNAME} -p={PASSWORD} -a={APARTMENT_ID} -f={FLOOR_ID} -d={DORMITORY_NUMBER} -A={AREA_ID} -s={SEND_KEY} -t={THRESHOLD}
```

其中: 

- USERNAME: 北邮校园卡账号
- PASSWORD: 北邮校园卡密码
- APARTMENT_ID: 公寓ID
    - 西土城#10: d5cf9743f0864692a18c25efb02bf16a
- FLOOR_ID: 楼层，如一层为"1"
- DORMITORY_NUMBER: 宿舍号，如学10的101为"10-101"
- AREA_ID: 校区，西土城校区为"1"，沙河校区为"2"
- SEND_KEY: server酱的发送密钥
- THRESHOLD: 阈值，只有全部电量低于阈值时才通过server酱发送信息，默认值为0(阈值限制不启用)
