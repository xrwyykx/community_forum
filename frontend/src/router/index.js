import { createRouter, createWebHistory } from 'vue-router'
import { ElMessage } from 'element-plus'

const routes = [
  {
    path: '/login',
    name: 'Login',
    component: () => import('../views/Login.vue'),
    meta: { requiresAuth: false }
  },
  {
    path: '/register',
    name: 'Register',
    component: () => import('../views/Register.vue'),
    meta: { requiresAuth: false }
  },
  {
    path: '/profile',
    name: 'Profile',
    component: () => import('../views/Profile.vue'),
    meta: { requiresAuth: true }
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

// 路由守卫
router.beforeEach((to, from, next) => {
  // 检查cookie中是否存在session
  const cookies = document.cookie.split(';');
  const sessionCookie = cookies.find(cookie => cookie.trim().startsWith('SESSION='));
  const hasSession = !!sessionCookie;
  
  console.log('当前路由:', to.path);
  console.log('是否需要认证:', to.meta.requiresAuth);
  console.log('是否有session:', hasSession);
  console.log('当前所有cookies:', document.cookie);
  console.log('解析后的cookies:', cookies);
  if (sessionCookie) {
    console.log('找到session cookie:', sessionCookie.trim());
  }
  
  // 如果需要认证但没有session，重定向到登录页
  if (to.meta.requiresAuth && !hasSession) {
    console.log('未登录，重定向到登录页');
    ElMessage.warning('请先登录');
    next('/login');
  } else if (to.path === '/login' && hasSession) {
    // 如果已登录但访问登录页，重定向到个人资料页
    console.log('已登录，重定向到个人资料页');
    next('/profile');
  } else {
    console.log('允许访问');
    next();
  }
});

export default router 