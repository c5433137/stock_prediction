import * as echarts from '../../ec-canvas/echarts';
import {http} from '../../utils/util.js';
const app = getApp();
let stockIndex = 0;
let stockInfo = {
  stockCode: '002156',
  stockIv: '1.0',
  stockPv: '2.0',
  stockOv: '10000.0',
  stockType: '0',
};
let chart={};
var option = {
  title: {
    text: '股票代码',
    left: 'center',
   subtext : '纯属虚构',
  subtextStyle:{
      fontWeight: 'bloder'
    }

  },
  legend: {
    data: ['实际股价', '预测值','回测数据'],
    top: 50,
    left: 'center',
    backgroundColor: '#999',
    z: 100
  },
  grid: {
    containLabel: true
  },
  tooltip: {
    show: true,
    trigger: 'axis'
  },
  xAxis: {
    type: 'category',
    boundaryGap: false,
    data: ['周一', '周二', '周三', '周四', '周五', '周六', '周日'],
    // show: false
  },
  yAxis: {
    x: 'center',
    type: 'value',
    splitLine: {
      lineStyle: {
        type: 'dashed'
      }
    }
    // show: false
  },
  series: [{
    name: '实际股价',
    type: 'line',
    smooth: true,
    data: []
  }, {
    name: '预测值',
    type: 'line',
    smooth: true,
    data: []
    }, {
      name: '回测数据',
      type: 'line',
      smooth: true,
      data: []
    }]
};


function initChart(canvas, width, height, dpr) {
  
  chart = echarts.init(canvas, null, {
    width: width,
    height: height,
    devicePixelRatio: dpr // new
  });
  canvas.setChart(chart);

  getData().then(res=>{
    option.xAxis.data = res.X
    option.series[0].data = res.Real[stockIndex]
    option.series[1].data = [null,...res.Predict[stockIndex]]
    // option.series[2].data = res.Ori[stockIndex]
    let v = res.NewValue[stockIndex].toFixed(2)
    console.log("v=",v)
    option.title.text = `股票代码:${res.Stock}`
    option.title.subtext = `下一个交易日预测价：${res.NewValue[stockIndex].toFixed(2)}`
    chart.setOption(option);
  })

  return chart;

  

  
}

function getData(){
  
  return new Promise((resolve,reject)=>{
    const { stockCode='',
      stockIv='',
      stockPv='',
      stockOv = '' } = stockInfo
      
    http({
      url: `/v1/get_stock_prediction_data?stockId=${stockCode}&iv=${stockIv}&pv=${stockPv}&ov=${stockOv}`
    }).then(
      res => {
    // debugger

          resolve(res.data.data)
      }, res => {
        reject()
      })
  })
}




Page({
  data: {
    ec: {
      onInit: initChart
    },
    stockList:[
      {
        'fullName':'开盘价',
        'value':'0'
      },
      {
        'fullName': '收盘价',
        'value': '1'
      },
      {
        'fullName': '当日最低价',
        'value': '2'
      },
      {
        'fullName': '当日最高价',
        'value': '3'
      },
    ],
    stockInfo:{
      stockCode: '002156',
      stockIv: '1.0',
      stockPv: '2.0',
      stockOv: '10000.0',
      stockType:'0',
      stockTypeName:'开盘价'
    },
    predictPrice:''
  },
  stockSelect(e) {
    var index = e.detail.value
    var stockInfo = this.data.stockList[index]
    this.setData({
      "stockInfo.stockTypeName": stockInfo.fullName,
     
    }, () => {
      stockIndex= stockInfo.value
      this.searchBind()
    })
  },
  getTargetVal(e) {
    let name = e.target.dataset.name;
    this.setData({
      [name]: e.detail.value
    },()=>{
      stockInfo = Object.assign(stockInfo, this.data.stockInfo)

    })
    

  },
  searchBind(){
    stockInfo = Object.assign(stockInfo, this.data.stockInfo)
   
    getData().then(res => {
      option.xAxis.data = res.X
      option.series[0].data = res.Real[stockIndex]
      option.series[1].data = [null, ...res.Predict[stockIndex]]
      option.title.text = `股票代码:${res.Stock}`
      option.title.subtext = `下一个交易日预测价：${res.NewValue[stockIndex].toFixed(2)}`
      
      // todo:这里没有获取到chart实例
      console.log("chart=", chart)
      chart.setOption(option)
    })


  },

  onReady() {
  }
});
