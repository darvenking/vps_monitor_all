<template>
  <div class="page-layout">
    <header-top></header-top>
    <div class="container">
      <side-bar class="left"></side-bar>
      <div class="content">
        <router-view></router-view>
      </div>
    </div>
    <c-footer/>
  </div>
</template>
<script setup lang="ts">
import HeaderTop from '/@/views/Layout/HeaderTop.vue'
import CFooter from '/@/views/Layout/CFooter.vue'
import SideBar from '/@/views/Layout/SideBar.vue'
</script>
<style lang="scss" scoped>
.page-layout {
  height: 100vh;
}
.container {
  height: calc(100% - 100px);
  display: flex;
  background: #e9e9e9;
  .left {
    width: 150px;
  }
  .content {
    flex: 1;
    margin: 10px;
    padding: 20px;
    background: #fff;
    box-sizing: border-box;
    overflow-y: auto;
  }
}
</style>
