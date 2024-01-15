<style scoped>

</style>


<template>
    <div class=" flex flex-col text-amber-50 h-screen">
        <Header class=" flex h-[40px] w-full items-center bg-gray-900/50" v-show="state.toggleHidden"></Header>
        <div>
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
import { reactive } from 'vue';


const state :{
  toggleHidden: boolean
} = reactive({
  toggleHidden: true,
})
const windowToggleAndTop = () => {
  state.toggleHidden = !state.toggleHidden
  WindowSetAlwaysOnTop(!state.toggleHidden)
}

</script>