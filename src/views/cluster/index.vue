<template>
  <div class="app-container">
    <el-form ref="form" :model="form" label-width="150px">
      <el-form-item label="User Name">
        <el-input v-model="form.userName" placeholder="填写应用所属的用户名" style="width: 200px"/>
      </el-form-item>
      <el-form-item label="App Id">
        <el-input v-model="form.appId" placeholder="填写应用ID" style="width: 200px"/>
      </el-form-item>
      <el-form-item label="Queue Name">
        <el-select
          v-model="form.queueName"
          filterable
          placeholder="选择队列名称"
          style="width: 500px"
        >
          <el-option v-for="item in data" :key="item.toString()" :label="item.toString()" :value="item.toString()" />
        </el-select>
      </el-form-item>
      <el-form-item label="Perfom Random q?">
<!--        <div><el-col :span="2" class="line">disable option</el-col></div>-->
        <el-switch v-model="form.fc" />
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
        userName: '',
        appId: '',
        queueName: '',
        date1: '',
        date2: '',
        fc: false
      },
      data: '',
      code: ''
    }
  },
  mounted() {
    this.fetchAllQ()
  },
  methods: {
    onSubmit() {
      this.$message('submit!')
      this.$post('/balance',
        this.form)
        .then((resp) => {
          this.$message({
            message: resp.data.toString(),
            type: resp.code === 0 ? 'success' : 'error'
          })
        })
    },
    onCancel() {
      this.$message({
        message: 'cancel!',
        type: 'warning'
      })
    },
    fetchAllQ() {
      this.$get('/qqq', '').then((response) => {
        console.log(response)
        this.data = response.data
      })
    },
    onsuccess(response) {
      this.$message(response.data)
    }
  }
}
</script>

<style scoped>
.line{
  text-align: center;
}
</style>

