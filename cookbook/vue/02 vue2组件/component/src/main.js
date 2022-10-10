import Vue from 'vue'
import App from './App'
import MyUI from '../modules/my-ui'
import "../modules/my-ui/common.css"

// 调用 `MyUI.install(Vue)`
// 全局注册
// Vue.use(MyUI)
// 按需注册
Vue.use(MyUI,{
  comments:[
    'MyButton'
  ]
})

Vue.config.productionTip = false

new Vue({
  render: h => h(App),
}).$mount('#app')
