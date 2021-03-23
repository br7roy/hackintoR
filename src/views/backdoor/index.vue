<template>
  <div style="padding:30px;">
    <el-alert :closable="false" title="没有把握的情况下不要在此页面做任何操作" type="success" />
    <el-form ref="form" :model="form" >
    <el-form-item label="输入 secreat" >
      <el-input
          placeholder="secreat..."
          v-model="form.secreat"
          clearable>
      </el-input>
    </el-form-item>
      <el-form-item label="输入 command">
      <el-input
        type="textarea"
        placeholder="echo -e ..."
        v-model="form.cmd"
        clearable>
      </el-input>
    </el-form-item>
    <el-form-item>
      <el-button type="primary" @click="onSubmit">Execute!</el-button>
    </el-form-item>
    </el-form>
  </div>
</template>
<script>

export default {
  data() {
    return {
      form: {
        cmd: ''
      },
      data: [],
      code: '',
      ckbox: false,
      baseURL: this.baseURL
    }
  },
  mounted() {
  },
  created() {
    // 为渲染,较之挂载,块那么一点
  },
  methods: {
    onSubmit() {
      this.$message('As your command!')
      this.$post('/localshell',
        this.form)
        .then((response) => {
          this.$message({
            message: response.data.toString(),
            type: response.code === 0 ? 'success' : 'error'
          })
        })
    },
    onsuccess(response) {
      this.$message(response.data)
    }
  }
}
</script>
