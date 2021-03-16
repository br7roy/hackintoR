<template>
  <div style="padding:30px;">
    <el-alert :closable="false" title="上生新所: 支持批量excel上传,上传以后自动化解析并且上传至客户可读取的位置" type="success" />
    <el-upload
      class="upload-demo"
      drag="true"
      action="ssupload"
      multiple
    :on-error="onerrorr"
    :on-success="onsuccess">
      <i class="el-icon-upload"></i>
      <div class="el-upload__text">将文件拖到此处，或<em>点击上传</em></div>
      <div class="el-upload__tip" slot="tip">只能上传jpg/png文件，且不超过500kb</div>
    </el-upload>
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
      ckbox: false,
      baseURL: this.baseURL
    }
  },
  mounted() {
    this.queryHxtp()
  },
  created() {
    alert(this.baseURL)
    // 为渲染,较之挂载,块那么一点
  },
  methods: {
    onSubmit() {
      this.$message('submit!')
      console.log(JSON.stringify(this.form))
      this.$post('/definePoly',
        this.form)
        .then((response) => {
          this.$message({
            message: response.data.toString(),
            type: response.code === 0 ? 'success' : 'error'
          })
        })
    },
    onsuccess(response){
      this.$message(response.data)
    },
    onerrorr(err, file, fileList) {
      this.$message({
        message: '失败啦',
        type: 'error'
      })
    }
  }
}
</script>
