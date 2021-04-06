<template>
  <div class="app-container">
    <el-alert :closable="false" title="批量修改job下所有task状态,启停还未实现" type="warning" />
    <Sch/>
    <el-form ref="form" :model="form" label-width="120px">
      <el-form-item :rules="{require:true}" label="从这个JobId开始">
        <el-col :span="15">
          <el-input-number v-model="form.jobId1" placeholder="填入JobId" style="width: 300px;"/>
        </el-col>
      </el-form-item>
      -
      <el-form-item :rules="{require:true}" label="至这个JobId结束">
        <el-col :span="15">
          <el-input-number v-model="form.jobId2" placeholder="填入JobId" style="width: 300px"/>
        </el-col>
      </el-form-item>
      <el-form-item :rules="{require:true}" label="Ok or fail">
        <el-switch v-model="form.stat" />
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="onSubmit">提交</el-button>
        <el-button>取消</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
import Sch from '@/views/sch'
export default {
  components: {
    Sch
  },
  data() {
    return {
      form: {
        jobIds: [],
        stat: false,
        jobId1: 0,
        jobId2: 0
      },
      data: '',
      code: ''
    }
  },
  methods: {
    onSubmit() {
      for (let i = this.form.jobId1; i <= this.form.jobId2; i++) {
        this.form.jobIds.push(i)
      }
      if (!this.form.stat) {
        this.form.stat = 5
      } else {
        this.form.stat = 4
      }
      this
        .$post('bj', this.form)
        .then(response => {
          if (response.code === 0) {
            this.$message({
              message: response.data.toString(),
              type: response.code === 0 ? 'success' : 'error'
            })
          }
        })
        .catch(function(error) {
          console.log(error)
        })
    }
  }
}
</script>

