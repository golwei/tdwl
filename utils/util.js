const formatTime = date => {
  const year = date.getFullYear()
  const month = date.getMonth() + 1
  const day = date.getDate()
  const hour = date.getHours()
  const minute = date.getMinutes()
  const second = date.getSeconds()

  return [year, month, day].map(formatNumber).join('/') + ' ' + [hour, minute, second].map(formatNumber).join(':')
}

const formatNumber = n => {
  n = n.toString()
  return n[1] ? n : '0' + n
}


const getData =
  (url) => {
    return new Promise(
      (resolve, reject) => {
        wx.request({
          url: url, //仅为示例，并非真实的接口地址
          data: {
            x: '',
            y: ''
          },
          header: {
            'content-type': 'application/json' // 默认值
          },
          success: function (res) {
            resolve(res.data)
            // console.log(url)
            // page.setData({
            //   res: res.data
            // })
            // console.log(res.data)
          }
        })
      }
    )
  }

module.exports = {
  formatTime: formatTime,
  getData: getData
}

