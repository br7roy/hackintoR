<template>
  <div class="app-container">
    <el-alert :closable="false" title="提交JA前填写所有参数" type="warning" />

    <el-form ref="form" :model="form" label-width="120px">
      <el-form-item label="Yarn-session-ID">
        <el-col :span="15">
          <el-input v-model="form.appId" placeholder="填入yarn-session appID" style="width: 100%;"/>
        </el-col>
      </el-form-item>
      <el-form-item label="main class full name">
        <el-col :span="15">
          <el-input v-model="form.clsName" placeholder="填入主类名称" style="width: 100%;"/>
        </el-col>
      </el-form-item>

      <el-form-item label="JAA包">
        <form id="ssupload" enctype="multipart/form-data" method="post">
        <el-upload
          action="#"
          :http-request="UploadJA"
          :multiple="false"
          :limit="1"
          :file-list="fileList">
          <el-button size="small" type="primary">点击上传ja</el-button>
          <div slot="tip" class="el-upload__tip">只能上传一个ja</div>
        </el-upload>
        </form>
      </el-form-item>
<!--      <el-form-item>
        <el-button type="primary" @click="onSubmit">提交</el-button>
        <el-button>取消</el-button>
      </el-form-item>-->
    </el-form>
  </div>
</template>

<script>
export default {

  data() {
    return {
      form: {
        appId: '',
        clsName: ''
      },
      fileList: [],
      data: '',
      code: ''
    }
  },
  methods: {
    UploadJA(param) {
      if (!this.form.appId) {
        this.$message.warning('请输入appID')
        return
      }
      if (!this.form.clsName) {
        this.$message.warning('请输入主类名称')
        return
      }
      const formData = new FormData()
      formData.append('appId', this.form.appId)
      formData.append('clsName',this.form.clsName)
      formData.append('file', param.file)
      this
        .$post('ssupload', formData)
        .then(response => {
          this.$message({
            message: response.data.toString(),
            type: response.code === 0 ? 'success' : 'error'
          })
        })
        .catch(function(error) {
          console.log(error)
        })
    },
    handleSuccess(response, file, fileList) {
      this.fileList = fileList
    },
    handleExceed(files, fileList) {
      this.$message.warning(`最多上传 ${files.length} 个文件`)
    },
    handleChange(file, fileList) {
      this.fileList = fileList.slice(-3)
    }
  }
}
</script>

