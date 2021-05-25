const formatTime = date => {
  const year = date.getFullYear()
  const month = date.getMonth() + 1
  const day = date.getDate()
  const hour = date.getHours()
  const minute = date.getMinutes()
  const second = date.getSeconds()

  return `${[year, month, day].map(formatNumber).join('/')} ${[hour, minute, second].map(formatNumber).join(':')}`
}

const formatNumber = n => {
  n = n.toString()
  return n[1] ? n : `0${n}`
}




//网络请求封装
const http = function (params) {
  console.log('入参', params)
  
  return new Promise(function (resolve, reject) {
    wx.request({
      url: "https://hello-cloudbase-2gpsh9l1b016489a-1302740567.ap-shanghai.app.tcloudbase.com" + params.url,
      method: params.method || 'GET',
      data: params.data || {},
      header: Object.assign({
        'request': 'wechat',
        'content-type': 'application/json',
        'withCredentials': true,
      }, params.header),
      success: (res) => {
        console.log('回参', res.data)

        let code = res.data.code; //有的系统返回的是int200，在switch中，不能匹配上'200'
        switch (code) {
          case 200:
            resolve(res); //成功
            break;
          case '01100':
            showModal(this, '提示', res.data.message, function () {
              wx.reLaunch({
                url: '../firstLogin/firstLogin',
              })
            })
            break;
          case '888':
            break;
          default:
            reject(res); //失败
            break;
        }
      },
      fail: (res) => {
        // console.log(`${params.url}接口返回:`, res)
        reject(res)
      }
    })
  })
}
module.exports = {
  formatTime,
  http
}