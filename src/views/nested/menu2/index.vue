<template>
  <div style="padding:30px;">
    <el-alert :closable="false" title="定点地理围栏跑参数: 目前支持(日)任务" type="success" />
    <el-form ref="form" :model="form" label-width="120px">
      <el-form-item label="开始时间">
        <el-col :span="15">
          <el-date-picker v-model="form.startTime" type="date" placeholder="Pick a date" style="width: 100%;" />
        </el-col>
      </el-form-item>
      <el-form-item label="结束时间">
        <el-col :span="15">
          <el-date-picker v-model="form.endTime" type="date" placeholder="Pick a date" style="width: 100%;" />
        </el-col>
      </el-form-item>
      <el-form-item label="任务类型">

<!--        <el-checkbox-group v-model="form.types">-->
<!--          <el-checkbox-->
<!--              v-for="(item,index) in data"-->
<!--              :key="index"-->
<!--              value="item.Type"-->
<!--              >画像名:{{ item.Name }} | 类型编号:{{item.Type}}</el-checkbox>-->
<!--        </el-checkbox-group>-->
        <ul>
          <li v-for="(item,index) in data" :key="index">
            <el-checkbox   :key="index+1" @change="changeBox(item,index,$event)">
              画像名:{{ item.Name }} | 类型编号:{{item.Type}}
            </el-checkbox>
          </li>
        </ul>

<!--        <div  v-for="(item,index) in data" style="width: 25%;float: left">-->
<!--          <el-checkbox v-model="data[index].Name"-->
<!--        </div>-->

<!--        <el-select v-model="form.queueName" placeholder="选择任务类型">-->
<!--          <el-option v-for="item in data" :key="item.Type" :label="'画像名:'+item.Name+'-&#45;&#45;类型编号:'+item.Type" :value="item.value" />-->
<!--        </el-select>-->
      </el-form-item>
      <el-form-item
        v-for="(uuid, index) in form.uuids"
        :key="uuid.key"
        :label="'UUID' + index"
        :prop="'uuids.' + index + '.value'"
        :rules="{
          required: true, message: 'UUIDbu能为空,并且是=18位的数字', trigger: 'blur',pattern: /^\d{18}$/
        }"
      >
        <el-input v-model="uuid.value" /><el-button @click.prevent="removeDomain(uuid)">删除</el-button>
      </el-form-item>
      <el-form-item>
        <el-button @click="addDomain">新增UUID</el-button>
        <el-button @click="resetForm('form')">重置</el-button>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="onSubmit">Perform</el-button>
        <el-button @click="onCancel">Cancel</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>
<script>

export default {
  data() {
    return {
      form: {
        startTime: '',
        endTime: '',
        uuids: [{
          value: ''
        }],
        types: []
      },
      data: [],
      code: '',
      ckbox: false
    }
  },
  mounted() {
    this.queryHxtp()
  },
  created() {
    // 为渲染,较之挂载,块那么一点
  },
  methods: {
    onSubmit() {
      this.$message('submit!')
      this.$post('/definePoly',
        this.form)
        .then((response) => {
          this.$message({
            message: response.data.toString(),
            type: response.code === 0 ? 'success' : 'error'
          })
        })
    },
    onCancel() {
      this.$message({
        message: 'cancel!',
        type: 'warning'
      })
    },
    resetForm(formName) {
      this.$refs[formName].resetFields()
    },
    removeDomain(item) {
      var index = this.form.uuids.indexOf(item)
      if (index !== -1) {
        this.form.uuids.splice(index, 1)
      }
    },
    addDomain() {
      this.form.uuids.push({
        value: '',
        key: Date.now()
      })
    },
    queryHxtp() {
      this.$post('/hxtp',this.form)
        .then((response) => {
          this.data = response.data
        }).catch(err =>{
          console.log(err)
        })
    },
    changeBox(item,index,event){
      if (event == true) {
        this.form.types.push(item.Type)
      } else {
        if (index !== -1) {
          this.form.types.splice(index, 1)
        }
      }
    }
  }
}
</script>
