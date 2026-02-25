
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="活动名称:" prop="name">
    <el-input v-model="formData.name" :clearable="false" placeholder="请输入活动名称" />
</el-form-item>
        <el-form-item label="活动类型:" prop="type">
    <el-select v-model="formData.type" placeholder="请选择活动类型" style="width:100%" filterable :clearable="false">
       <el-option v-for="item in ['线上活动','线下活动','混合活动']" :key="item" :label="item" :value="item" />
    </el-select>
</el-form-item>
        <el-form-item label="活动内容:" prop="content">
    <RichEdit v-model="formData.content"/>
</el-form-item>
        <el-form-item label="开始时间:" prop="startTime">
    <el-date-picker v-model="formData.startTime" type="date" style="width:100%" placeholder="选择日期" :clearable="false" />
</el-form-item>
        <el-form-item label="结束时间:" prop="endTime">
    <el-date-picker v-model="formData.endTime" type="date" style="width:100%" placeholder="选择日期" :clearable="false" />
</el-form-item>
        <el-form-item label="活动状态:" prop="status">
    <el-select v-model="formData.status" placeholder="请选择活动状态" style="width:100%" filterable :clearable="false">
       <el-option v-for="item in ['未开始','进行中','已结束','已取消']" :key="item" :label="item" :value="item" />
    </el-select>
</el-form-item>
        <el-form-item label="活动图片:" prop="images">
    <SelectImage
     multiple
     v-model="formData.images"
     file-type="image"
     />
</el-form-item>
        <el-form-item label="联系方式:" prop="contact">
    <el-input v-model="formData.contact" :clearable="false" placeholder="请输入联系方式" />
</el-form-item>
        <el-form-item label="活动地点:" prop="location">
    <el-input v-model="formData.location" :clearable="false" placeholder="请输入活动地点" />
</el-form-item>
        <el-form-item label="活动预算:" prop="budget">
    <el-input-number v-model="formData.budget" style="width:100%" :precision="2" :clearable="false" />
</el-form-item>
        <el-form-item label="组织者:" prop="organizer">
    <el-input v-model="formData.organizer" :clearable="false" placeholder="请输入组织者" />
</el-form-item>
        <el-form-item>
          <el-button :loading="btnLoading" type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
  createChwActivity,
  updateChwActivity,
  findChwActivity
} from '@/api/example/activityService'

defineOptions({
    name: 'ChwActivityForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
// 图片选择组件
import SelectImage from '@/components/selectImage/selectImage.vue'
// 富文本组件
import RichEdit from '@/components/richtext/rich-edit.vue'


const route = useRoute()
const router = useRouter()

// 提交按钮loading
const btnLoading = ref(false)

const type = ref('')
const formData = ref({
            name: '',
            type: null,
            content: '',
            startTime: new Date(),
            endTime: new Date(),
            status: null,
            images: [],
            contact: '',
            location: '',
            budget: 0,
            organizer: '',
        })
// 验证规则
const rule = reactive({
               name : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               type : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               content : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               startTime : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               endTime : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               status : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               organizer : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               id : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findChwActivity({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
// 保存按钮
const save = async() => {
      btnLoading.value = true
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return btnLoading.value = false
            let res
           switch (type.value) {
             case 'create':
               res = await createChwActivity(formData.value)
               break
             case 'update':
               res = await updateChwActivity(formData.value)
               break
             default:
               res = await createChwActivity(formData.value)
               break
           }
           btnLoading.value = false
           if (res.code === 0) {
             ElMessage({
               type: 'success',
               message: '创建/更改成功'
             })
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
