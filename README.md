# stock_prediction
股票预测
#### 简介
frond_end:目录是微信小程序代码。
其余为golang的后端代码。
日k数据采用急速api，免费额度每日限制请求100次。

本项目使用ekf来进行股价预测，通过小程序输入代码（目前仅支持sh、sz市场），可调参数为iv\pv\ov，
分别代表初始误差、模型误差、ov误差。ekf中对股价模型采用简单的线性模型，预测结果主要依赖于
历史采集的日k数据。

#### 小程序内容

![小程序内容](https://i.loli.net/2021/05/23/HMK3uWjqy7wOo1g.png)


#### 微信小程序体验二维码

![临时体验二维码](https://i.loli.net/2021/05/23/qjoGNtTM8BV5lsw.png)



