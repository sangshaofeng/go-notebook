<template>
  <div class="header">
    <div class="content-wrapper">
      <Menu mode="horizontal" active-name="1">
        <h4>文档管理</h4>
        <Submenu name="3" class="avatar-wrapper">
            <template slot="title">
              <img class="avatar" src="../assets/logo.png" alt="logo">
            </template>
            <MenuItem name="3-1">
              <router-link to="">个人主页</router-link>
            </MenuItem>
            <MenuItem name="3-1">
              <router-link to="">新建文档</router-link>
            </MenuItem>
            <MenuItem name="3-1">
              <router-link to="/login">登出</router-link>
            </MenuItem>
        </Submenu>
      </Menu>
    </div>
  </div>
</template>

<script>
  export default {
    
  }
</script>

<style lang="less" scoped>
a {
  color: #515a6e;
}
.header {
  width: 100%;
  height: 50px;
  text-align: left;
  background-color: #fff;
  box-shadow: 0 1px 3px rgba(26,26,26,.1);
  position: fixed;
  top: 0;
  left: 0;
  z-index: 98;
  .content-wrapper {
    width: 1000px;
    height: 100%;
    margin: 0 auto;
    .ivu-menu-horizontal {
      height: 50px;
      line-height: 50px;
      border-bottom: none;
    }
    .ivu-menu-horizontal.ivu-menu-light:after {
      background: #fff;
    }
    h4 {
      display: inline-block;
    }
    .avatar-wrapper {
      float: right;
      .avatar {
        width: 30px;
        vertical-align: middle;
      }
    }

  }
}
</style>