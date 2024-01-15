<style scoped>


</style>
<!-- 头部  -->
<template>
    <div class=" flex px-2 py-2">
        <div>{{ state.appName }}</div>
        <div class=" flex-grow h-full" style="--wails-draggable:drag"></div>
        <div class=" flex space-x-2">
            <span class=" cursor-pointer flex items-center" @click="WindowHide()"><el-icon size="20px"><i-ep-ArrowDown /></el-icon></span>
            <span class=" cursor-pointer flex items-center" @click="WindowMinimise()"><el-icon size="20px"><i-ep-Minus /></el-icon></span>
            <span class=" cursor-pointer flex items-center" @click="WindowToggleMaximise()"><el-icon size="20px"><i-ep-FullScreen /></el-icon></span>
            <span class=" cursor-pointer flex items-center" @click="Quit()"><el-icon size="20px"><i-ep-Close /></el-icon></span>
        </div>
    </div>

</template>
<script setup lang="ts">
import {WindowHide, WindowMinimise ,WindowToggleMaximise,Quit } from '../../wailsjs/runtime'
import { reactive, ref } from 'vue'
import { GetAppName } from '../../wailsjs/go/mainApp/MainApp'

const state: {
    appName: string
} = reactive({
    appName: ''
})


GetAppName().then((name) => {
    state.appName = name;
})


</script>
