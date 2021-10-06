import Vue from 'vue'
import App from './App.vue'

Vue.config.productionTip = false

import '@fortawesome/fontawesome-free/css/all.css'
import '@fortawesome/fontawesome-free/js/all.js'

import TaskList from "./components/TaskList.vue"
import TaskItem from "./components/TaskItem.vue"

Vue.component("task-list", TaskList)
Vue.component("task-item", TaskItem)

new Vue({
  render: h => h(App),
}).$mount('#app')
