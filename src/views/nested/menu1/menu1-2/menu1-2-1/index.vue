<template>
  <div style="padding:30px;">
    <el-alert :closable="false" title="daily注意:" type="success" />
    <el-form ref="form" :model="form" label-width="120px">
      <el-form-item label="开始时间">
        <el-col :span="15">
          <el-date-picker v-model="form.startTime" type="date" placeholder="Pick a date" style="width: 100%;"/>
        </el-col>
      </el-form-item>
      <el-form-item label="结束时间">
        <el-col :span="15">
          <el-date-picker v-model="form.endTime" type="date" placeholder="Pick a date" style="width: 100%;"/>
        </el-col>
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
        endTime: ''
      },
      data: '',
      code: ''
    }
  },
  mounted() {
    // this.fetchAllQ()
  },
  created() {
    // 为渲染,较之挂载,块那么一点
  },
  methods: {
    onSubmit() {
      this.$message('submit!')
      this.$post('/daily',
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
    }
  }
}
</script>

<style scoped>
.line{
  text-align: center;
}
</style>

