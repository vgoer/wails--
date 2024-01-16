<style scoped>

</style>


<template>
    <div class=" flex flex-col text-amber-50 h-screen">
        <Header class=" flex h-[40px] w-full items-center bg-gray-900/50" v-show="state.toggleHidden"></Header>
        <div class=" flex justify-center items-center h-full">
              <el-button type="primary" @click="useHanld()" class=" bg-[#E94234] font-bold" color="#E94234">点我</el-button>
        </div>
        <div class="relative  text-end">
            <el-icon 
              class="text-cyan-50/50 hover:text-cyan-50 hover:cursor-pointer absolute right-1 bottom-1.5"
              size="30px"
              @click="windowToggleAndTop"
              ><i-ep-ChromeFilled /></el-icon>
        </div>
    </div>
</template>

<script lang="ts" setup>
import Header from './components/Header.vue'

import { WindowSetAlwaysOnTop } from '../wailsjs/runtime'
import { GetNotify } from '../wailsjs/go/mainApp/MainApp'
import { ElMessage } from 'element-plus'
import { reactive } from 'vue';


const state :{
  toggleHidden: boolean
  title: string
} = reactive({
  toggleHidden: true,
  title: ''
})
const windowToggleAndTop = () => {
  state.toggleHidden = !state.toggleHidden
  WindowSetAlwaysOnTop(!state.toggleHidden)
}


const useHanld = () => {
    ElMessage({
      message: '你个大傻春',
      type: 'error',
    })
}



</script>