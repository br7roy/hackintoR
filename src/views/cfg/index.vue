<template>
  <div style="padding:30px;">
    <el-alert :closable="false" title="没有把握的情况下不要在此页面做任何操作" type="success"/>
    <el-form ref="form" :model="form">
      <el-form-item label="输入 secreat">
        <el-input
          v-model="form.RemoteAddr"
          placeholder="secreat..."
          clearable
        />
      </el-form-item>
      <el-form-item label="输入 command">
        <el-input
          v-model="form.cmd"
          type="textarea"
          placeholder="echo -e ..."
          clearable
        />
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
        App: {
          RemoteAddr: '',
          RemotePort: '',
          UsingNC: false,
          UsingSH: false,
          Test: false,
          Schedule: {
            UserName: ''
          },
          ScpInfo: {
            Root: '',
            DataSync: '',
            GpsConfigure: '',
            BatchDaily: '',
            BatchResult: '',
            BatchPy: '',
            Month: ''
          },
          Bus: {
            Portrait: null
          },
          Db: {
            UserName: '',
            Password: '',
            Url: '',
            Port: 0,
            StoreHouse: '',
            Type: ''
          }
        },
        Queue: []
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
      this.$message('update config!')
      this.$post('/bumpconf',
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
    },
    fetchAllQ() {
      this.$get('/readconf', '').then((response) => {
        console.log(response)
        this.data = response.data
      })
    }
  }
}
</script>
