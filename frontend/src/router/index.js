import Vue from 'vue'
import Router from 'vue-router'
import Home from '@/view/Home'
import ChatRoom from '@/view/ChatRoom'
Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'Home',
      component: Home
    },
      {
      path: '/chatroom',
      name: 'ChatRoom',
      component: ChatRoom
    }
  ]
})
