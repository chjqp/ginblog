import { createRouter, createWebHashHistory } from 'vue-router';

// 使用路由懒加载
const Login = () => import(/* webpackChunkName: "login" */ '../views/Login.vue');
const Admin = () => import(/* webpackChunkName: "admin" */ '../views/Admin.vue');

const routes = [
  {
    path: '/',
    redirect: '/login' // 根路径重定向到登录页面
  },
  {
    path: '/login',
    name: 'login',
    component: Login
  },
  {
    path: '/admin',
    name: 'admin',
    component: Admin
  }
];

const router = createRouter({
  history: createWebHashHistory(),
  routes
});

export default router;
