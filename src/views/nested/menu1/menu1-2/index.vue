<template>
  <div style="padding:50px;margin-right: fill">
    <el-alert :closable="false" title="商场客流" />
    <el-form ref="form" :model="form" label-width="150px">
      <el-form-item label="产线执行?">
        <el-switch
          v-model="form.prdEnv"
          active-color="#ff4949"
          inactive-color="#13ce66"
          @change="change(form.prdEnv)"
        />
      </el-form-item>
    </el-form>
<!--    <el-col :span="3">产线执行?</el-col>-->
<!--    <el-switch-->
<!--      v-model="form.prdEnv"-->
<!--      active-color="#ff4949"-->
<!--      inactive-color="#13ce66"-->
<!--      @change="change(form.prdEnv)"-->
<!--    />-->
    <router-view />
  </div>

</template>
<script>
import { mapGetters } from 'vuex'

export default {

  data() {
    return {
      form: {
        prdEnv: false
      }
    }
  },
  computed: {
    ...mapGetters([
      'prdEnv'
    ])
  },
  methods: {
    change(env) {
      this.$store.dispatch('app/changeEnv', env)
    }
  }

}

</script>
