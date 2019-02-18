<template>
  <div class="home">
    <Header/>
    <div class="main-content-wrapper">
      <div class="left-wrapper">
        <div class="head">
          <h4>文档目录</h4>
          <Button type="text" icon="md-add">新建</Button>
        </div>
        <ul class="cate-list">
          <li>
            <Icon type="md-folder" />
            <span>前端</span>
          </li>
          <li>
            <Icon type="md-folder" />
            <span>前端2017-10-10</span>
          </li>
          <li>
            <Icon type="md-folder" />
            <span>后端</span>
          </li>
          <li>
            <Icon type="md-folder" />
            <span>后端2017-10-10</span>
          </li>
        </ul>
      </div>
      <div class="right-wrapper">
        <Spin fix></Spin>
      </div>
    </div>
  </div>
</template>

<script>
import Header from '../components/Header'
export default {
  components: {
    Header
  },
  data () {
    return {
      
    }
  }
}
</script>

<style lang="less" scoped>
.ivu-spin-fix {
  background: transparent;
}
.home {
  .main-content-wrapper {
    width: 1000px;
    height: auto;
    padding: 10px 0;
    margin: 60px auto;
    text-align: left;
    display: flex;
    flex-flow: row nowrap;
    justify-content: flex-start;
    align-items: flex-start;
    .left-wrapper {
      display: inline-block;
      width: 280px;
      min-height: 280px;
      background: #fff;
      padding: 10px;
      border-radius: 4px;
      box-shadow: 0 1px 1px 0 rgba(0,0,0,.1);
      position: relative;
      .head {
        height: 34px;
        line-height: 34px;
        border-bottom: 1px solid #f8f8f9;
        display: flex;
        flex-flow: row nowrap;
        justify-content: space-between;
        align-items: center;
      }
      .cate-list {
        margin-top: 12px;
        max-height: 400px;
        overflow: auto;
        li {
          font-size: 14px;
          height: 28px;
          line-height: 28px;
          cursor: pointer;
          transition: all 0.3s ease-in-out;
          &:hover {
            color: #2d8cf0;
          }
          span {
            margin-left: 8px;
          }
        }
      }
    }
    .right-wrapper {
      display: inline-block;
      margin-left: 20px;
      width: 700px;
      min-height: 180px;
      background: #fff;
      padding: 10px;
      border-radius: 4px;
      box-shadow: 0 1px 1px 0 rgba(0,0,0,.1);
      position: relative;
    }
  }
}
</style>

