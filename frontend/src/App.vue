<template>
  <div class="app-container">
    <router-view></router-view>
  </div>
</template>

<script setup>
import { onMounted } from 'vue'
import { useRouter } from 'vue-router'

const router = useRouter()

onMounted(() => {
  // 如果没有token且不在登录或注册页面，重定向到登录页
  const token = localStorage.getItem('token')
  const currentPath = router.currentRoute.value.path
  if (!token && currentPath !== '/login' && currentPath !== '/register') {
    router.push('/login')
  }
})
</script>

<style>
.app-container {
  min-height: 100vh;
  background-color: #f5f7fa;
}

/* 全局样式重置 */
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

body {
  font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, 'Helvetica Neue', Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}
</style> 