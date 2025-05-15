<template>
  <div class="profile-container">
    <el-row :gutter="20">
      <el-col :span="6">
        <el-card class="user-card">
          <div class="user-avatar">
            <el-avatar :size="100" :src="userInfo.avatar || defaultAvatar" />
            <h3 class="username">{{ userInfo.username }}</h3>
          </div>
          <div class="user-stats">
            <div class="stat-item">
              <span class="stat-value">{{ userInfo.posts || 0 }}</span>
              <span class="stat-label">发帖</span>
            </div>
            <div class="stat-item">
              <span class="stat-value">{{ userInfo.followers || 0 }}</span>
              <span class="stat-label">粉丝</span>
            </div>
            <div class="stat-item">
              <span class="stat-value">{{ userInfo.following || 0 }}</span>
              <span class="stat-label">关注</span>
            </div>
          </div>
          <div class="button-group">
            <el-button type="primary" @click="showEditDialog" class="edit-button">
              编辑资料
            </el-button>
            <el-button type="warning" @click="handleLogout" class="logout-button">
              退出登录
            </el-button>
          </div>
        </el-card>
      </el-col>
      
      <el-col :span="18">
        <el-card>
          <template #header>
            <div class="card-header">
              <span>个人资料</span>
            </div>
          </template>
          <el-descriptions :column="1" border>
            <el-descriptions-item label="用户名">{{ userInfo.username }}</el-descriptions-item>
            <el-descriptions-item label="邮箱">{{ userInfo.email }}</el-descriptions-item>
            <el-descriptions-item label="注册时间">{{ userInfo.createTime }}</el-descriptions-item>
            <el-descriptions-item label="个人简介">{{ userInfo.bio || '暂无简介' }}</el-descriptions-item>
          </el-descriptions>
        </el-card>

        <el-card class="mt-20">
          <template #header>
            <div class="card-header">
              <span>我的帖子</span>
            </div>
          </template>
          <el-table :data="userPosts" style="width: 100%">
            <el-table-column prop="title" label="标题" />
            <el-table-column prop="createTime" label="发布时间" width="180" />
            <el-table-column prop="views" label="浏览" width="100" />
            <el-table-column prop="likes" label="点赞" width="100" />
            <el-table-column label="操作" width="150">
              <template #default="scope">
                <el-button size="small" @click="viewPost(scope.row)">查看</el-button>
                <el-button size="small" type="danger" @click="deletePost(scope.row)">删除</el-button>
              </template>
            </el-table-column>
          </el-table>
        </el-card>
      </el-col>
    </el-row>

    <!-- 编辑资料对话框 -->
    <el-dialog v-model="dialogVisible" title="编辑个人资料" width="500px">
      <el-form :model="editForm" :rules="rules" ref="editFormRef" label-width="100px">
        <el-form-item label="头像">
          <el-upload
            class="avatar-uploader"
            action="/api/common/upload"
            :show-file-list="false"
            :on-success="handleAvatarSuccess"
          >
            <img v-if="editForm.avatar" :src="editForm.avatar" class="avatar" />
            <el-icon v-else class="avatar-uploader-icon"><Plus /></el-icon>
          </el-upload>
        </el-form-item>
        <el-form-item label="用户名" prop="username">
          <el-input v-model="editForm.username" />
        </el-form-item>
        <el-form-item label="邮箱" prop="email">
          <el-input v-model="editForm.email" />
        </el-form-item>
        <el-form-item label="个人简介" prop="bio">
          <el-input type="textarea" v-model="editForm.bio" :rows="4" />
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="dialogVisible = false">取消</el-button>
          <el-button type="primary" @click="handleUpdateProfile" :loading="loading">
            确认
          </el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { Plus } from '@element-plus/icons-vue'
import axios from 'axios'
import { useRouter } from 'vue-router'

const router = useRouter()
const defaultAvatar = 'https://cube.elemecdn.com/3/7c/3ea6beec64369c2642b92c6726f1epng.png'
const userInfo = ref({})
const userPosts = ref([])
const dialogVisible = ref(false)
const loading = ref(false)
const editFormRef = ref(null)

const editForm = reactive({
  username: '',
  email: '',
  bio: '',
  avatar: ''
})

const rules = {
  username: [
    { required: true, message: '请输入用户名', trigger: 'blur' },
    { min: 3, max: 20, message: '长度在 3 到 20 个字符', trigger: 'blur' }
  ],
  email: [
    { required: true, message: '请输入邮箱地址', trigger: 'blur' },
    { type: 'email', message: '请输入正确的邮箱地址', trigger: 'blur' }
  ]
}

const fetchUserInfo = async () => {
  try {
    const response = await axios.post('/api/common/get-user-info')
    if (response.data.code === 0) {
      userInfo.value = response.data.data
      Object.assign(editForm, response.data.data)
    } else if (response.data.code === 40000) {
      // 未登录或会话过期
      ElMessage.warning('请先登录')
      router.push('/login')
    }
  } catch (error) {
    console.error('获取用户信息失败:', error)
    if (error.response && error.response.status === 401) {
      ElMessage.warning('请先登录')
      router.push('/login')
    } else {
      ElMessage.error('获取用户信息失败')
    }
  }
}

const fetchUserPosts = async () => {
  try {
    const response = await axios.post('/api/user/get-user-posts', {
      page: 1,
      pageSize: 10
    })
    if (response.data.code === 0) {
      userPosts.value = response.data.data.posts
    } else if (response.data.code === 40000) {
      // 未登录或会话过期
      ElMessage.warning('请先登录')
      router.push('/login')
    }
  } catch (error) {
    console.error('获取帖子列表失败:', error)
    if (error.response && error.response.status === 401) {
      ElMessage.warning('请先登录')
      router.push('/login')
    } else {
      ElMessage.error('获取帖子列表失败')
    }
  }
}

const showEditDialog = () => {
  dialogVisible.value = true
}

const handleAvatarSuccess = (response) => {
  if (response.code === 0) {
    editForm.avatar = response.data.url
  } else {
    ElMessage.error('上传头像失败')
  }
}

const handleUpdateProfile = async () => {
  if (!editFormRef.value) return
  
  await editFormRef.value.validate(async (valid) => {
    if (valid) {
      loading.value = true
      try {
        const response = await axios.post('/api/user/update-user-info', editForm)
        if (response.data.code === 0) {
          ElMessage.success('更新成功')
          dialogVisible.value = false
          await fetchUserInfo()
        } else {
          ElMessage.error(response.data.message || '更新失败')
        }
      } catch (error) {
        ElMessage.error('更新失败，请稍后重试')
      } finally {
        loading.value = false
      }
    }
  })
}

const viewPost = (post) => {
  // 实现查看帖子详情
}

const deletePost = async (post) => {
  try {
    await ElMessageBox.confirm('确定要删除这篇帖子吗？', '提示', {
      type: 'warning'
    })
    const response = await axios.delete('/api/user/delete-post', {
      data: { postId: post.id }
    })
    if (response.data.code === 0) {
      ElMessage.success('删除成功')
      await fetchUserPosts()
    } else {
      ElMessage.error(response.data.message || '删除失败')
    }
  } catch (error) {
    if (error !== 'cancel') {
      ElMessage.error('删除失败，请稍后重试')
    }
  }
}

const handleLogout = async () => {
  try {
    await ElMessageBox.confirm('确定要退出登录吗？', '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    })
    
    // 清除 cookie
    document.cookie = 'SESSION=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;'
    
    ElMessage.success('已退出登录')
    router.push('/login')
  } catch (error) {
    if (error !== 'cancel') {
      ElMessage.error('退出登录失败')
    }
  }
}

onMounted(async () => {
  try {
    await fetchUserInfo()
    await fetchUserPosts()
  } catch (error) {
    console.error('初始化数据失败:', error)
  }
})
</script>

<style scoped>
.profile-container {
  padding: 20px;
}

.user-card {
  text-align: center;
}

.user-avatar {
  padding: 20px 0;
}

.username {
  margin: 10px 0;
  color: #303133;
}

.user-stats {
  display: flex;
  justify-content: space-around;
  margin: 20px 0;
  padding: 10px 0;
  border-top: 1px solid #ebeef5;
  border-bottom: 1px solid #ebeef5;
}

.stat-item {
  display: flex;
  flex-direction: column;
}

.stat-value {
  font-size: 20px;
  font-weight: bold;
  color: #303133;
}

.stat-label {
  font-size: 14px;
  color: #909399;
}

.button-group {
  display: flex;
  flex-direction: column;
  gap: 10px;
  margin-top: 20px;
}

.edit-button,
.logout-button {
  width: 100%;
}

.logout-button {
  margin-top: 10px;
}

.mt-20 {
  margin-top: 20px;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.avatar-uploader {
  text-align: center;
}

.avatar-uploader .avatar {
  width: 100px;
  height: 100px;
  border-radius: 50%;
}

.avatar-uploader .el-upload {
  border: 1px dashed #d9d9d9;
  border-radius: 50%;
  cursor: pointer;
  position: relative;
  overflow: hidden;
}

.avatar-uploader .el-upload:hover {
  border-color: #409eff;
}

.avatar-uploader-icon {
  font-size: 28px;
  color: #8c939d;
  width: 100px;
  height: 100px;
  line-height: 100px;
  text-align: center;
}
</style> 