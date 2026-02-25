
<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" @keyup.enter="onSubmit">
            <el-form-item label="活动ID" prop="id">
  <el-input v-model.number="searchInfo.id" placeholder="搜索条件" />
</el-form-item>
            

        <template v-if="showAllQuery">
          <!-- 将需要控制显示状态的查询条件添加到此范围内 -->
        </template>

        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
          <el-button link type="primary" icon="arrow-down" @click="showAllQuery=true" v-if="!showAllQuery">展开</el-button>
          <el-button link type="primary" icon="arrow-up" @click="showAllQuery=false" v-else>收起</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
        <div class="gva-btn-list">
            <el-button  type="primary" icon="plus" @click="openDialog()">新增</el-button>
            <el-button  icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="onDelete">删除</el-button>
            <ExportTemplate  template-id="example_ChwActivity" />
            <ExportExcel  template-id="example_ChwActivity" filterDeleted/>
            <ImportExcel  template-id="example_ChwActivity" @on-success="getTableData" />
        </div>
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="id"
        @selection-change="handleSelectionChange"
        @sort-change="sortChange"
        >
        <el-table-column type="selection" width="55" />
        
            <el-table-column align="left" label="活动名称" prop="name" width="120" />

            <el-table-column align="left" label="活动类型" prop="type" width="120" />

            <el-table-column sortable align="left" label="开始时间" prop="startTime" width="180">
   <template #default="scope">{{ formatDate(scope.row.startTime) }}</template>
</el-table-column>
            <el-table-column sortable align="left" label="结束时间" prop="endTime" width="180">
   <template #default="scope">{{ formatDate(scope.row.endTime) }}</template>
</el-table-column>
            <el-table-column sortable align="left" label="活动状态" prop="status" width="120" />

            <el-table-column align="left" label="参与人数" prop="participants" width="120" />

            <el-table-column align="left" label="联系方式" prop="contact" width="120" />

            <el-table-column align="left" label="活动地点" prop="location" width="120" />

            <el-table-column align="left" label="活动预算" prop="budget" width="120" />

            <el-table-column align="left" label="组织者" prop="organizer" width="120" />

            <el-table-column sortable align="left" label="活动ID" prop="id" width="120" />

        <el-table-column align="left" label="操作" fixed="right" :min-width="appStore.operateMinWith">
            <template #default="scope">
            <el-button  type="primary" link class="table-button" @click="getDetails(scope.row)"><el-icon style="margin-right: 5px"><InfoFilled /></el-icon>查看</el-button>
            <el-button  type="primary" link icon="edit" class="table-button" @click="updateChwActivityFunc(scope.row)">编辑</el-button>
            <el-button   type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
            </template>
        </el-table-column>
        </el-table>
        <div class="gva-pagination">
            <el-pagination
            layout="total, sizes, prev, pager, next, jumper"
            :current-page="page"
            :page-size="pageSize"
            :page-sizes="[10, 30, 50, 100]"
            :total="total"
            @current-change="handleCurrentChange"
            @size-change="handleSizeChange"
            />
        </div>
    </div>
    <el-drawer destroy-on-close :size="appStore.drawerSize" v-model="dialogFormVisible" :show-close="false" :before-close="closeDialog">
       <template #header>
              <div class="flex justify-between items-center">
                <span class="text-lg">{{type==='create'?'新增':'编辑'}}</span>
                <div>
                  <el-button :loading="btnLoading" type="primary" @click="enterDialog">确 定</el-button>
                  <el-button @click="closeDialog">取 消</el-button>
                </div>
              </div>
            </template>

          <el-form :model="formData" label-position="top" ref="elFormRef" :rules="rule" label-width="80px">
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
          </el-form>
    </el-drawer>

    <el-drawer destroy-on-close :size="appStore.drawerSize" v-model="detailShow" :show-close="true" :before-close="closeDetailShow" title="查看">
            <el-descriptions :column="1" border>
                    <el-descriptions-item label="活动名称">
    {{ detailForm.name }}
</el-descriptions-item>
                    <el-descriptions-item label="活动类型">
    {{ detailForm.type }}
</el-descriptions-item>
                    <el-descriptions-item label="活动内容">
    <RichView v-model="detailForm.content" />
</el-descriptions-item>
                    <el-descriptions-item label="开始时间">
    {{ detailForm.startTime }}
</el-descriptions-item>
                    <el-descriptions-item label="结束时间">
    {{ detailForm.endTime }}
</el-descriptions-item>
                    <el-descriptions-item label="活动状态">
    {{ detailForm.status }}
</el-descriptions-item>
                    <el-descriptions-item label="参与人数">
    {{ detailForm.participants }}
</el-descriptions-item>
                    <el-descriptions-item label="活动图片">
    <el-image style="width: 50px; height: 50px; margin-right: 10px" :preview-src-list="returnArrImg(detailForm.images)" :initial-index="index" v-for="(item,index) in detailForm.images" :key="index" :src="getUrl(item)" fit="cover" />
</el-descriptions-item>
                    <el-descriptions-item label="联系方式">
    {{ detailForm.contact }}
</el-descriptions-item>
                    <el-descriptions-item label="活动地点">
    {{ detailForm.location }}
</el-descriptions-item>
                    <el-descriptions-item label="活动预算">
    {{ detailForm.budget }}
</el-descriptions-item>
                    <el-descriptions-item label="组织者">
    {{ detailForm.organizer }}
</el-descriptions-item>
                    <el-descriptions-item label="活动ID">
    {{ detailForm.id }}
</el-descriptions-item>
            </el-descriptions>
        </el-drawer>

  </div>
</template>

<script setup>
import {
  createChwActivity,
  deleteChwActivity,
  deleteChwActivityByIds,
  updateChwActivity,
  findChwActivity,
  getChwActivityList
} from '@/api/example/activityService'
import { getUrl } from '@/utils/image'
// 图片选择组件
import SelectImage from '@/components/selectImage/selectImage.vue'
// 富文本组件
import RichEdit from '@/components/richtext/rich-edit.vue'
import RichView from '@/components/richtext/rich-view.vue'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict ,filterDataSource, returnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'
import { useAppStore } from "@/pinia"

// 导出组件
import ExportExcel from '@/components/exportExcel/exportExcel.vue'
// 导入组件
import ImportExcel from '@/components/exportExcel/importExcel.vue'
// 导出模板组件
import ExportTemplate from '@/components/exportExcel/exportTemplate.vue'


defineOptions({
    name: 'ChwActivity'
})

// 提交按钮loading
const btnLoading = ref(false)
const appStore = useAppStore()

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

// 自动化生成的字典（可能为空）以及字段
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
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
               type : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               content : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               startTime : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               endTime : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               status : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               organizer : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
})

const elFormRef = ref()
const elSearchFormRef = ref()

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})
// 排序
const sortChange = ({ prop, order }) => {
  const sortMap = {
    CreatedAt:"created_at",
    ID:"id",
            startTime: 'start_time',
            endTime: 'end_time',
            status: 'status',
            id: 'id',
  }

  let sort = sortMap[prop]
  if(!sort){
   sort = prop.replace(/[A-Z]/g, match => `_${match.toLowerCase()}`)
  }

  searchInfo.value.sort = sort
  searchInfo.value.order = order
  getTableData()
}
// 重置
const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

// 搜索
const onSubmit = () => {
  elSearchFormRef.value?.validate(async(valid) => {
    if (!valid) return
    page.value = 1
    getTableData()
  })
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  const table = await getChwActivityList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () =>{
}

// 获取需要的字典 可能为空 按需保留
setOptions()


// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
    multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
    ElMessageBox.confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    }).then(() => {
            deleteChwActivityFunc(row)
        })
    }

// 多选删除
const onDelete = async() => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async() => {
      const ids = []
      if (multipleSelection.value.length === 0) {
        ElMessage({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      multipleSelection.value &&
        multipleSelection.value.map(item => {
          ids.push(item.id)
        })
      const res = await deleteChwActivityByIds({ ids })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功'
        })
        if (tableData.value.length === ids.length && page.value > 1) {
          page.value--
        }
        getTableData()
      }
      })
    }

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateChwActivityFunc = async(row) => {
    const res = await findChwActivity({ id: row.id })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteChwActivityFunc = async (row) => {
    const res = await deleteChwActivity({ id: row.id })
    if (res.code === 0) {
        ElMessage({
                type: 'success',
                message: '删除成功'
            })
            if (tableData.value.length === 1 && page.value > 1) {
            page.value--
        }
        getTableData()
    }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)

// 打开弹窗
const openDialog = () => {
    type.value = 'create'
    dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
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
        }
}
// 弹窗确定
const enterDialog = async () => {
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
                closeDialog()
                getTableData()
              }
      })
}

const detailForm = ref({})

// 查看详情控制标记
const detailShow = ref(false)


// 打开详情弹窗
const openDetailShow = () => {
  detailShow.value = true
}


// 打开详情
const getDetails = async (row) => {
  // 打开弹窗
  const res = await findChwActivity({ id: row.id })
  if (res.code === 0) {
    detailForm.value = res.data
    openDetailShow()
  }
}


// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  detailForm.value = {}
}


</script>

<style>

</style>
