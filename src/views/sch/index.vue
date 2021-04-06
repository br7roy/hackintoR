<template>
  <div style="padding: initial">
    <el-alert :closable="false" title="调度增强" />
    <el-form ref="form" :model="form" label-width="150px">
      <el-form-item :rules="{require:true}" label="scheduleName">
        <el-input
          v-model="form.schUserName"
          type="text"
          placeholder="不带邮件域名的用户名"
          active-color="#ff4949"
          inactive-color="#13ce66"
          style="width: 200px"
          @change="change(form.schUserName)"
        />
      </el-form-item>
    </el-form>
    <router-view />
  </div>
</template>
<script>
import { mapGetters } from 'vuex'

export default {
  data() {
    return {
      form: {
        schUserName: ''
      }
    }
  },
  computed: {
    ...mapGetters([
      'schUserName'
    ])
  },
  methods: {
    change(schUserName) {
      console.log('invoked')
      this.$store.dispatch('app/changeSchUserName', schUserName)
    }
  }
}
</script>
