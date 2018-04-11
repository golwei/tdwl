//logs.js
const util = require('../../utils/util.js')
const url =
  "this is url"

Page({
  data: {
    goods: [],
    logs: []
  },
  onLoad: function () {
    console.log("1")
    util.getData(this, url)
    console.log("1")
    util.getData(this, url)
    console.log("1")
    util.getData(this, url)
    console.log("1")
    util.getData(this, url)
    console.log("1")
    util.getData(this, url)
    console.log("1")
    util.getData(this, url)
    console.log("1")
    util.getData(this, url)
    console.log("1")
    util.getData(this, url)
    console.log("1")
    // this.myfunc()
    this.setData({
      logs: (wx.getStorageSync('logs') || []).map(log => {
        return util.formatTime(new Date(log))
      })
    })
  },
  myfunc: function () {
    wx.request({
      url: 'http://wcqt.bid/admin/goods.json',
      success: res => {
        console.log(res)
      }
    })
  },
})


