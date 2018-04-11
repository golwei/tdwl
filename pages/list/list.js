// pages/list/list.js
const util = require('../../utils/util.js')
const getCategories = "http://wcqt.bid/admin/categories.json"
const getProducts = "http://wcqt.bid/admin/products.json"
Page({

  /**
   * 页面的初始数据
   */
  data: {
    categories: [],
    products: []
  },

  /**
   * 生命周期函数--监听页面加载
   */
  onLoad: function (options) {
    util.getData(getCategories)
      .then(res => {
        this.setData({
          categories: res
        })
      })
    util.getData(getProducts)
      .then(res => {
        this.setData({
          products: res
        })
      })
  },

  /**
   * 生命周期函数--监听页面初次渲染完成
   */
  onReady: function () {

  },

  /**
   * 生命周期函数--监听页面显示
   */
  onShow: function () {

  },

  /**
   * 生命周期函数--监听页面隐藏
   */
  onHide: function () {

  },

  /**
   * 生命周期函数--监听页面卸载
   */
  onUnload: function () {

  },

  /**
   * 页面相关事件处理函数--监听用户下拉动作
   */
  onPullDownRefresh: function () {

  },

  /**
   * 页面上拉触底事件的处理函数
   */
  onReachBottom: function () {

  },

  /**
   * 用户点击右上角分享
   */
  onShareAppMessage: function () {

  }
})